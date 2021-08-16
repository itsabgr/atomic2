package atomic2

import (
	"sync/atomic"
)

type Uint64 uint64

func (n *Uint64) Add(increase uint64) uint64 {
	return atomic.AddUint64((*uint64)(n), increase)
}

func (n *Uint64) Sub(decrease uint64) uint64 {
	return atomic.AddUint64((*uint64)(n), ^(decrease - 1))
}

func (n *Uint64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64((*uint64)(n), new)
}
func (n *Uint64) CAS(test, value uint64) bool {
	return atomic.CompareAndSwapUint64((*uint64)(n), test, value)
}
func (n *Uint64) Set(value uint64) {
	atomic.StoreUint64((*uint64)(n), value)
}
func (n *Uint64) Get() uint64 {
	return atomic.LoadUint64((*uint64)(n))
}
