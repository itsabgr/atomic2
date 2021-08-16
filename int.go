package atomic2

import (
	"sync/atomic"
)

type Int uintptr

func (n *Int) Add(increase int) uintptr {
	return atomic.AddUintptr((*uintptr)(n), uintptr(increase))
}

func (n *Int) Sub(decrease int) int {
	return int(atomic.AddUintptr((*uintptr)(n), ^(uintptr(decrease) - 1)))
}

func (n *Int) Swap(new int) (old int) {
	return int(atomic.SwapUintptr((*uintptr)(n), uintptr(new)))
}
func (n *Int) CAS(test, value int) bool {
	return atomic.CompareAndSwapUintptr((*uintptr)(n), uintptr(test), uintptr(value))
}
func (n *Int) Set(value int) {
	atomic.StoreUintptr((*uintptr)(n), uintptr(value))
}
func (n *Int) Get() int {
	return int(atomic.LoadUintptr((*uintptr)(n)))
}
