package once

import (
	"sync"
	"sync/atomic"
)

// Once is backwards compatible re-implementation of sync.Once.
// See https://golang.org/pkg/sync/#Once
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do is a backwards compatible re-implementation of Do from sync.Once
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// Error is similar to Once, except it returns an error value.
type Error struct {
	m    sync.Mutex
	done uint32
	err  error
}

// Do runs the specified function only once, but all callers gets the same
// result from that single execution.
func (o *Error) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return o.err
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		o.err = f()
	}
	return o.err
}

// Value is similar to Once, except it returns a value.
type Value[T any] struct {
	m     sync.Mutex
	done  uint32
	value T
}

// Do runs the specified function only once, but all callers gets the same
// result from that single execution.
func (o *Value[T]) Do(f func() T) T {
	if atomic.LoadUint32(&o.done) == 1 {
		return o.value
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		o.value = f()
	}
	return o.value
}

// ValueError is similar to Once, except it return a (value, error) tuple
type ValueError[T any] struct {
	m     sync.Mutex
	done  uint32
	value T
	err   error
}

// Do runs the specified function only once, but all callers gets the same
// result from that single execution.
func (o *ValueError[T]) Do(f func() (T, error)) (T, error) {
	if atomic.LoadUint32(&o.done) == 1 {
		return o.value, o.err
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		o.value, o.err = f()
	}
	return o.value, o.err
}
