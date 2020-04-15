package errorx

import (
"fmt"
"path"
"runtime"
)

type XError struct {
	Location string
	Msg      string
	Err      error
}

func (x *XError) Error() string {
	if x == nil {
		return ""
	}
	if x.Err == nil {
		return x.Msg + x.Location
	} else {
		return x.Msg + x.Location + " >> " + x.Err.Error()
	}
}

func ErrorMsg(msg string) error {
	if msg == "" {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)
	return &XError{
		Msg:      msg,
		Location: printMark(file, line),
		Err:      nil,
	}
}

func Error(err error) error {
	if err == nil {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)
	return &XError{
		Location: printMark(file, line),
		Err:      err,
	}
}

func ErrorWithMsg(err error, msg string) error {
	if err == nil {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)
	return &XError{
		Msg:      msg,
		Location: printMark(file, line),
		Err:      err,
	}
}

func printMark(file string, line int) string {
	return fmt.Sprintf(" -%s:%d", path.Base(file), line)
}
