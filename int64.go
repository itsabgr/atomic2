package atomic2

import (
	"sync/atomic"
)

type Int64 int64

func (n *Int64) Add(increase int64) int64 {
	return atomic.AddInt64((*int64)(n), increase)
}

func (n *Int64) Sub(decrease int64) int64 {
	return atomic.AddInt64((*int64)(n), ^(decrease - 1))
}

func (n *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64((*int64)(n), new)
}
func (n *Int64) CAS(test, value int64) bool {
	return atomic.CompareAndSwapInt64((*int64)(n), test, value)
}
func (n *Int64) Set(value int64) {
	atomic.StoreInt64((*int64)(n), value)
}
func (n *Int64) Get() int64 {
	return atomic.LoadInt64((*int64)(n))
}
