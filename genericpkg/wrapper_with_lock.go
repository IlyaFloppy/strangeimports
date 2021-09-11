package genericpkg

import (
	"fmt"
	"sync"
)

type WrapperWithLock[T any] interface {
	PrintWithLock()
}

func NewWrapperWithLock[T any](value T) WrapperWithLock[T] {
	return &wrapperWithLock[T]{
		Object: value,
	}
}

type wrapperWithLock[T any] struct {
	Lock   sync.Mutex
	Object T
}

func (w *wrapperWithLock[T]) PrintWithLock() {
	w.Lock.Lock()
	defer w.Lock.Unlock()

	fmt.Println(w.Object)
}
