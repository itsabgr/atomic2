package atomic2

import (
	"sync/atomic"
	"unsafe"
)

type Ptr struct {
	Ptr unsafe.Pointer
}

func (ptr Ptr) Set(another unsafe.Pointer) {
	atomic.StorePointer(&ptr.Ptr, another)
}
func (ptr Ptr) Get() (another unsafe.Pointer) {
	return atomic.LoadPointer(&ptr.Ptr)
}
func (ptr Ptr) Swap(new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(&ptr.Ptr, new)
}
func (ptr Ptr) CAS(test, new unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(&ptr.Ptr, test, new)
}
func (ptr Ptr) SwapValue(value interface{}) (old interface{}) {
	return *(*interface{})(atomic.SwapPointer(&ptr.Ptr, unsafe.Pointer(&value)))
}
func (ptr Ptr) SetValue(value interface{}) {
	ptr.Set(unsafe.Pointer(&value))
}
func (ptr Ptr) GetValue() interface{} {
	return *(*interface{})(ptr.Get())
}
