package logx

type Unit int64

func (u Unit) CalB(b int64) int64 {
	return b * int64(u)
}

const (
	_       = iota
	KB Unit = 1 << (iota * 10)
	MB
	GB
	TB
)
