package randx

import (
	"net"
	"time"
)

func Int63() int64 { r := GetRand(); defer r.Release(); return r.Int63() }

func Uint32() uint32 { r := GetRand(); defer r.Release(); return r.Uint32() }

func Uint64() uint64 { r := GetRand(); defer r.Release(); return r.Uint64() }

func Int31() int32 { r := GetRand(); defer r.Release(); return r.Int31() }

func Int() int { r := GetRand(); defer r.Release(); return r.Int() }

func Int63n(n int64) int64 { r := GetRand(); defer r.Release(); return r.Int63n(n) }

func Int31n(n int32) int32 { r := GetRand(); defer r.Release(); return r.Int31n(n) }

func Intn(n int) int { r := GetRand(); defer r.Release(); return r.Intn(n) }

func Float64() float64 { r := GetRand(); defer r.Release(); return r.Float64() }

func Float32() float32 { r := GetRand(); defer r.Release(); return r.Float32() }

func Perm(n int) []int { r := GetRand(); defer r.Release(); return r.Perm(n) }

func Shuffle(n int, swap func(i, j int)) { r := GetRand(); defer r.Release(); r.Shuffle(n, swap) }

func Read(p []byte) (n int, err error) { r := GetRand(); defer r.Release(); return r.Read(p) }

func NormFloat64() float64 { r := GetRand(); defer r.Release(); return r.NormFloat64() }

func ExpFloat64() float64 { r := GetRand(); defer r.Release(); return r.ExpFloat64() }

func String(length int, flag int) string {
	r := GetRand()
	defer r.Release()
	return r.String(length, flag)
}

func StringLib(length int, str string) string {
	r := GetRand()
	defer r.Release()
	return r.StringLib(length, str)
}

func StringArray(list []string) string {
	r := GetRand()
	defer r.Release()
	return r.StringArray(list)
}

func RangeString(s string) string {
	r := GetRand()
	defer r.Release()
	return r.RangeString(s)
}

func Phone() string {
	r := GetRand()
	defer r.Release()
	return r.Phone()
}

// RandRange 范围随机 [min, max]
func RandRange(min int, max int) int {
	r := GetRand()
	defer r.Release()
	return r.RandRange(min, max)
}

// RandRangeInt32 范围随机 [min, max]
func RandRangeInt32(min int32, max int32) int {
	r := GetRand()
	defer r.Release()
	return r.RandRangeInt32(min, max)
}

// RandRangeInt64 范围随机 [min, max]
func RandRangeInt64(min int64, max int64) int64 {
	r := GetRand()
	defer r.Release()
	return r.RandRangeInt64(min, max)
}

func Bool() bool {
	r := GetRand()
	defer r.Release()
	return r.Bool()
}
func ChineseIP() net.IP {
	r := GetRand()
	defer r.Release()
	return r.ChineseIP()
}

func ChineseIPStr() string {
	r := GetRand()
	defer r.Release()
	return r.ChineseIPStr()
}

func Port() int {
	r := GetRand()
	defer r.Release()
	return r.Port()
}

func IPAddress() string {
	r := GetRand()
	defer r.Release()
	return r.IPAddress()
}

func ChineseName() string {
	r := GetRand()
	defer r.Release()
	return r.ChineseName()
}

func WeightStringInt64MapOnce(m map[string]int64) string {
	r := GetRand()
	defer r.Release()
	return r.WeightStringInt64MapOnce(m)
}

func WeightMapOnce(m map[interface{}]int64) interface{} {
	r := GetRand()
	defer r.Release()
	return r.WeightMapOnce(m)
}

func WeightSliceNoRepeat(slice interface{}, calcWeight func(i int) int64, count int) ([]int, error) {
	r := GetRand()
	defer r.Release()
	return r.WeightSliceNoRepeat(slice, calcWeight, count)
}

func TimeRange(min, max time.Duration) time.Duration {
	r := GetRand()
	defer r.Release()
	return r.TimeRange(min, max)
}

// IntIntervalMultipleNoRepeat int区间不重复随机
// 范围 [begin, end)
func IntIntervalMultipleNoRepeat(begin, end, count int) []int {
	r := GetRand()
	defer r.Release()
	return r.IntIntervalMultipleNoRepeat(begin, end, count)
}
