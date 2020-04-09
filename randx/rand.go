package randx

import (
	"math/rand"
	"sync"
	"time"
)

type Rand struct {
	*rand.Rand
}

var (
	pool = sync.Pool{
		New: func() interface{} {
			return &Rand{rand.New(rand.NewSource(time.Now().UnixNano()))}
		},
	}
)

func GetRand() *Rand {
	return pool.Get().(*Rand)
}

func NewRand() *Rand {
	return GetRand()
}

func (r *Rand) Release() {
	pool.Put(r)
}

// RandRange 范围随机 [min, max]
func (r *Rand) RandRange(min int, max int) int {
	if max <= min {
		if max == min {
			return min
		}
		return 0
	}
	return r.Intn(max-min+1) + min
}

// RandRangeInt32 范围随机 [min, max]
func (r *Rand) RandRangeInt32(min int32, max int32) int {
	return r.RandRange(int(min), int(max))
}

// RandRangeInt64 范围随机 [min, max]
func (r *Rand) RandRangeInt64(min int64, max int64) int64 {
	if max <= min {
		if max == min {
			return min
		}
		return 0
	}
	return r.Int63n(max-min+1) + min
}

func (r *Rand) Bool() bool {
	return r.Intn(2) == 0
}

// IntIntervalMultipleNoRepeat int区间不重复随机
// 范围 [begin, end)
func (r *Rand) IntIntervalMultipleNoRepeat(begin, end, count int) []int {
	n := end - begin
	if n <= 0 || n < count {
		return nil
	}

	g := getOrderlyIntSlice(n)
	defer g.Release()

	choice := make([]int, count)
	for ; count > 0; count-- {
		c := r.Intn(n)
		choice[count-1] = g[c]
		g[c] = g[n-1]
		n--
	}

	for i, v := range choice {
		g[v] = v
		choice[i] += begin
	}

	return choice
}
