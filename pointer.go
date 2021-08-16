package atomic2

import (
	"sync/atomic"
	"unsafe"
)

type Ptr struct {
	ptr unsafe.Pointer
}

func (ptr Ptr) Set(another unsafe.Pointer) {
	atomic.StorePointer(&ptr.ptr, another)
}
func (ptr Ptr) Get() (another unsafe.Pointer) {
	return atomic.LoadPointer(&ptr.ptr)
}
func (ptr Ptr) Swap(new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(&ptr.ptr, new)
}
func (ptr Ptr) CAS(test, new unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(&ptr.ptr, test, new)
}
func (ptr Ptr) SwapValue(value interface{}) (old interface{}) {
	return *(*interface{})(atomic.SwapPointer(&ptr.ptr, unsafe.Pointer(&value)))
}
func (ptr Ptr) SetValue(value interface{}) {
	ptr.Set(unsafe.Pointer(&value))
}
func (ptr Ptr) GetValue() interface{} {
	return *(*interface{})(ptr.Get())
}
