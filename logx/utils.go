package logx

import (
	"os"
	"time"
)

func isFileExit(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func minLevel(levels ...Level) Level {
	if len(levels) == 0 {
		return OFF
	}
	minLevel := OFF
	for _, level := range levels {
		if level < minLevel {
			minLevel = level
		}
	}
	return minLevel
}

func fastTimeLocalFormatByte(t time.Time) []byte {
	y, mo, d := t.Date()
	abs := t.Unix() + 3600*8
	abs %= 86400
	h := abs / 3600
	abs %= 3600
	m := abs / 60
	s := abs % 60
	return []byte{
		byte('0' + y/1000),
		byte('0' + y/100%10),
		byte('0' + y/10%10),
		byte('0' + y%10),
		byte('-'),
		byte('0' + mo/10),
		byte('0' + mo%10),
		byte('-'),
		byte('0' + d/10),
		byte('0' + d%10),
		byte(' '),
		byte('0' + h/10),
		byte('0' + h%10),
		byte(':'),
		byte('0' + m/10),
		byte('0' + m%10),
		byte(':'),
		byte('0' + s/10),
		byte('0' + s%10),
	}
}
