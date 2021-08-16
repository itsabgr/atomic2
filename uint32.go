package atomic2

import (
	"sync/atomic"
)

type Uint32 uint32

func (n *Uint32) Add(increase uint32) uint32 {
	return atomic.AddUint32((*uint32)(n), increase)
}

func (n *Uint32) Sub(decrease uint32) uint32 {
	return atomic.AddUint32((*uint32)(n), ^(decrease - 1))
}

func (n *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32((*uint32)(n), new)
}
func (n *Uint32) CAS(test, value uint32) bool {
	return atomic.CompareAndSwapUint32((*uint32)(n), test, value)
}
func (n *Uint32) Set(value uint32) {
	atomic.StoreUint32((*uint32)(n), value)
}
func (n *Uint32) Get() uint32 {
	return atomic.LoadUint32((*uint32)(n))
}
