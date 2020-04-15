package logx

type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}

type LevelHooks map[Level][]Hook

func (l LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		l[level] = append(l[level], hook)
	}
}

func (l LevelHooks) Adds(hooks []Hook) {
	for _, hook := range hooks {
		l.Add(hook)
	}
}

func (l LevelHooks) minLevel() Level {
	minLevel := OFF
	for level := range l {
		if level < minLevel {
			minLevel = level
		}
	}
	return minLevel
}

func (l LevelHooks) Fire(level Level, entry *Entry) error {
	for _, hook := range l[level] {
		if err := hook.Fire(entry); err != nil {
			return err
		}
	}

	return nil
}
