package logx

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_callerSkip = 5
)

type Logger struct {
	opt *option

	fileEntryChanLength int64
	fileEntryChan       chan *Entry

	file *File

	hooksMutex sync.Mutex
	hooks      LevelHooks

	minLevel Level

	entryPool  *sync.Pool
	bufferPool *sync.Pool
}

func NewLogger(opts ...Options) *Logger {
	opt := newDefaultOption()
	for _, v := range opts {
		v(opt)
	}

	lg := &Logger{
		opt:           opt,
		fileEntryChan: make(chan *Entry, opt.ChannelBuffLength),
		bufferPool:    &sync.Pool{New: func() interface{} { return new(bytes.Buffer) }},
		hooks:         LevelHooks{},
	}
	lg.entryPool = &sync.Pool{
		New: func() interface{} { return &Entry{Logger: lg} },
	}

	lg.hooks.Adds(opt.Hooks)
	lg.minLevel = minLevel(opt.ConsoleLevel, opt.FileLevel, lg.hooks.minLevel())

	lg.run()

	return lg
}

func (l *Logger) Close() {
	if l == nil {
		return
	}
	close(l.fileEntryChan)
	l.fileEntryChanLength = 0
	l.file.Close()
}

func (l *Logger) log(skip int, level Level, args []interface{}) {
	if l == nil {
		return
	}

	entry := l.entryPool.Get().(*Entry)
	entry.Time = time.Now()
	entry.CallerSkip = skip
	entry.Level = level
	entry.Args = args

	if level >= l.opt.ConsoleLevel && l.opt.ConsoleFormat != nil {
		fmt.Println(l.opt.ConsoleFormat(entry))
	}

	entry.fireHooks()

	if level >= l.opt.FileLevel && l.opt.FileFormat != nil {
		if atomic.LoadInt64(&l.fileEntryChanLength) < l.opt.ChannelBuffLength {
			atomic.AddInt64(&l.fileEntryChanLength, 1)
			l.fileEntryChan <- entry
		} else {
			fmt.Println("Log stack overflow, discard.")
		}
	}
}

func (l *Logger) Wait() {
	for {
		if l.fileEntryChanLength == 0 {
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func (l *Logger) isLowLevel(level Level) bool {
	return level < l.minLevel
}

func (l *Logger) logf(skip int, level Level, format string, args []interface{}) {
	l.log(skip+1, level, []interface{}{fmt.Sprintf(format, args...)})
}

func (l *Logger) Debug(args ...interface{}) {
	if l.isLowLevel(DEBUG) {
		return
	}
	l.log(_callerSkip, DEBUG, args)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.isLowLevel(DEBUG) {
		return
	}
	l.logf(_callerSkip, DEBUG, format, args)
}

func (l *Logger) Info(args ...interface{}) {
	if l.isLowLevel(INFO) {
		return
	}
	l.log(_callerSkip, INFO, args)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.isLowLevel(INFO) {
		return
	}
	l.logf(_callerSkip, INFO, format, args)
}

func (l *Logger) Warn(args ...interface{}) {
	if l.isLowLevel(WARN) {
		return
	}
	l.log(_callerSkip, WARN, args)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.isLowLevel(WARN) {
		return
	}
	l.logf(_callerSkip, WARN, format, args)
}

func (l *Logger) Error(args ...interface{}) {
	if l.isLowLevel(ERROR) {
		return
	}
	l.log(_callerSkip, ERROR, args)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.isLowLevel(ERROR) {
		return
	}
	l.logf(_callerSkip, ERROR, format, args)
}

func (l *Logger) LogfSkip(skip int, level Level, format string, args ...interface{}) {
	if l.isLowLevel(level) {
		return
	}
	l.logf(skip, level, format, args)
}

func (l *Logger) refreshMinLevel() {
	l.minLevel = minLevel(l.opt.ConsoleLevel, l.opt.FileLevel, l.hooks.minLevel())
}

func (l *Logger) SetFileLevel(level Level) {
	l.opt.FileLevel = level
	l.refreshMinLevel()
}

func (l *Logger) SetConsoleLevel(level Level) {
	l.opt.ConsoleLevel = level
	l.refreshMinLevel()
}

func (l *Logger) run() {
	go func() {
		for entry := range l.fileEntryChan {
			l.writeFile(entry)
			atomic.AddInt64(&l.fileEntryChanLength, -1)
			entry.release()
		}
	}()
}

func (l *Logger) writeFile(entry *Entry) {
	f, err := l.opt.FileSplit(l, l.file, entry)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.file = f

	err = l.file.write(l.opt.FileFormat(entry))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *Logger) SetFileSize(size Unit) {
	l.opt.FileSizeMax = size
}

func (l *Logger) SetFileCount(count int) {
	l.opt.FileCountMax = count
}

func (l *Logger) AddHook(hook Hook) {
	l.hooksMutex.Lock()
	defer l.hooksMutex.Unlock()
	l.hooks.Add(hook)
	l.refreshMinLevel()
}

func (l *Logger) releaseEntry(entry *Entry) {
	entry.release()
	l.entryPool.Put(entry)
}
