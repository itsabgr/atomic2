package atomic2

import (
	"sync/atomic"
)

type Int32 int32

func (n *Int32) Add(increase int32) int32 {
	return atomic.AddInt32((*int32)(n), increase)
}

func (n *Int32) Sub(decrease int32) int32 {
	return atomic.AddInt32((*int32)(n), ^(decrease - 1))
}

func (n *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32((*int32)(n), new)
}
func (n *Int32) CAS(test, value int32) bool {
	return atomic.CompareAndSwapInt32((*int32)(n), test, value)
}
func (n *Int32) Set(value int32) {
	atomic.StoreInt32((*int32)(n), value)
}
func (n *Int32) Get() int32 {
	return atomic.LoadInt32((*int32)(n))
}
