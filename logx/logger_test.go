package logx

import (
	"fmt"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	_ = os.RemoveAll("./log")
	InitDefaultLogger(WithFileName("tt"), WithFileDir("./log"), WithFileLevel(INFO), WithConsoleLevel(ALL))
	Debug("debug", 1)
	Info("info", 1)
	Warn("warn", 1)
	Error("error", 1)
	GetDefaultLogger().Info("log.info", 2)

	// hook
	GetDefaultLogger().AddHook(&THook{})
	Error("error", 123)

	Wait()
}

type THook struct {
}

func (t *THook) Levels() []Level {
	return []Level{ERROR}
}
func (t *THook) Fire(entry *Entry) error {
	fmt.Println("hook fire", string(entry.ArgsDefaultFormat()))
	return nil
}
