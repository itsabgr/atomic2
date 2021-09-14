package atomic2

import (
	"sync/atomic"
)

type Uintptr uintptr

type Uint = Uintptr

func (n *Uintptr) Add(increase uintptr) uintptr {
	return atomic.AddUintptr((*uintptr)(n), increase)
}

func (n *Uintptr) Sub(decrease uintptr) uintptr {
	return atomic.AddUintptr((*uintptr)(n), ^(decrease - 1))
}

func (n *Uintptr) Swap(new uintptr) (old uintptr) {
	return atomic.SwapUintptr((*uintptr)(n), new)
}
func (n *Uintptr) CAS(test, value uintptr) bool {
	return atomic.CompareAndSwapUintptr((*uintptr)(n), test, value)
}
func (n *Uintptr) Set(value uintptr) {
	atomic.StoreUintptr((*uintptr)(n), value)
}
func (n *Uintptr) Get() uintptr {
	return atomic.LoadUintptr((*uintptr)(n))
}
