package logx

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type option struct {
	ChannelBuffLength int64

	FileName     string
	FileLevel    Level
	FileDir      string
	FileSplit    func(*Logger, *File, *Entry) (*File, error)
	FileSizeMax  Unit
	FileCountMax int
	FileFormat   func(*Entry) []byte

	ConsoleLevel  Level
	ConsoleFormat func(*Entry) string

	Hooks []Hook
}

type Options func(*option)

func newDefaultOption() *option {
	fileName := path.Base(filepath.ToSlash(os.Args[0]))
	if strings.HasSuffix(strings.ToLower(fileName), ".exe") {
		fileName = fileName[:len(fileName)-4]
	}
	return &option{
		ChannelBuffLength: 10000,

		FileName:     fileName,
		FileLevel:    OFF,
		FileDir:      "./log",
		FileSplit:    DefaultFileSplitBySize,
		FileSizeMax:  MB * 50,
		FileCountMax: 10,
		FileFormat:   defaultFileFormatFunc,

		ConsoleLevel:  INFO,
		ConsoleFormat: defaultConsoleFormatFunc,
	}
}

func (o *option) GetFilePath(no int) string {
	if no == 0 {
		return fmt.Sprintf("%s/%s.log", o.FileDir, o.FileName)
	}
	return fmt.Sprintf("%s/%s_%d.log", o.FileDir, o.FileName, no)
}

func WithFileLevel(lvl Level) Options {
	return func(o *option) { o.FileLevel = lvl }
}

func WithConsoleLevel(lvl Level) Options {
	return func(o *option) { o.ConsoleLevel = lvl }
}

func WithConsole(isShowConsole bool) Options {
	if !isShowConsole {
		return func(o *option) { o.ConsoleLevel = OFF }
	}
	return func(o *option) {}
}

func WithFileDir(dir string) Options {
	return func(o *option) { o.FileDir = dir }
}

func WithFileName(fileName string) Options {
	return func(o *option) { o.FileName = fileName }
}

func WithFileSize(fileSize Unit) Options {
	return func(o *option) { o.FileSizeMax = fileSize }
}

func WithFileMaxCount(maxCount int) Options {
	return func(o *option) { o.FileCountMax = maxCount }
}

func WithFileSplit(f func(*Logger, *File, *Entry) (*File, error)) Options {
	return func(o *option) { o.FileSplit = f }
}

func WithFileFormatFunc(f func(f *Entry) []byte) Options {
	return func(o *option) { o.FileFormat = f }
}

func WithConsoleFormatFunc(f func(f *Entry) string) Options {
	return func(o *option) { o.ConsoleFormat = f }
}

func WithHooks(hooks []Hook) Options {
	return func(o *option) { o.Hooks = hooks }
}
