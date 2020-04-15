package logx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Logger *Logger

	CallerSkip  int
	CallerFrame *runtime.Frame

	Buffer *bytes.Buffer

	Time  time.Time
	Level Level
	Args  []interface{}
}

func (e *Entry) GetCaller() *runtime.Frame {
	if e.CallerFrame == nil {
		rpc := make([]uintptr, 1)
		n := runtime.Callers(e.CallerSkip, rpc[:])
		if n < 1 {
			return nil
		}
		frame, _ := runtime.CallersFrames(rpc).Next()
		e.CallerFrame = &frame
	}
	return e.CallerFrame
}

func (e *Entry) ArgsDefaultFormat() []byte {
	var buffer bytes.Buffer
	for i := 0; i < len(e.Args); i++ {
		s := fmt.Sprint(e.Args[i])
		if len(s) == 12 && strings.HasPrefix(s, "0x") {
			data, err := json.Marshal(e.Args[i])
			if err == nil {
				s = "&" + string(data)
			}
		}
		buffer.WriteString(s + " ")
	}
	return buffer.Bytes()
}

func (e *Entry) fireHooks() {
	e.Logger.hooksMutex.Lock()
	defer e.Logger.hooksMutex.Unlock()
	err := e.Logger.hooks.Fire(e.Level, e)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
	}
}

func (e *Entry) getBuffer() *bytes.Buffer {
	if e.Buffer == nil {
		e.Buffer = e.Logger.bufferPool.Get().(*bytes.Buffer)
	}
	e.Buffer.Reset()
	return e.Buffer
}

func (e *Entry) release() {
	e.CallerFrame = nil
	if e.Buffer != nil {
		e.Logger.bufferPool.Put(e.Buffer)
		e.Buffer = nil
	}
}

func defaultFormatFileName(fName string) string {
	return path.Base(fName)
}

func defaultConsoleFormatFunc(f *Entry) string {
	b := f.getBuffer()
	b.Write(fastTimeLocalFormatByte(f.Time))
	b.WriteString(" [")
	b.WriteString(f.Level.ConsoleColorString())
	b.WriteString("] ")
	b.Write(f.ArgsDefaultFormat())
	b.WriteString("-")
	caller := f.GetCaller()
	b.WriteString(defaultFormatFileName(caller.File))
	b.WriteString(":")
	b.WriteString(strconv.Itoa(caller.Line))
	return b.String()
}

func defaultFileFormatFunc(f *Entry) []byte {
	b := f.getBuffer()
	b.Write(fastTimeLocalFormatByte(f.Time))
	b.WriteString(" [")
	b.WriteString(f.Level.String())
	b.WriteString("] ")
	b.Write(f.ArgsDefaultFormat())
	b.WriteString("-")
	caller := f.GetCaller()
	b.WriteString(defaultFormatFileName(caller.File))
	b.WriteString(":")
	b.WriteString(strconv.Itoa(caller.Line))
	b.WriteString("\n")
	return b.Bytes()
}
