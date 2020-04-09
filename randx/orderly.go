package randx

import "sync"

type stOrderlyIntSlice []int

// _orderlyIntSlice 有序int切片池
var _orderlyIntSlice = sync.Pool{
	New: func() interface{} { return stOrderlyIntSlice([]int{}) },
}

func getOrderlyIntSlice(count int) stOrderlyIntSlice {
	g := _orderlyIntSlice.Get().(stOrderlyIntSlice)
	if len(g) >= count {
		return g
	}
	_orderlyIntSlice.Put(g)

	g = stOrderlyIntSlice(make([]int, count))
	for i := 0; i < count; i++ {
		g[i] = i
	}
	return g
}

func (o stOrderlyIntSlice) Release() {
	_orderlyIntSlice.Put(o)
}
