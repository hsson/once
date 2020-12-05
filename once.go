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

// OnceError is similar to Once, except it returns an error value.
type OnceError struct {
	m    sync.Mutex
	done uint32
	err  error
}

// Do runs the specified function only once, but all callers gets the same
// result from that single execution.
func (o *OnceError) Do(f func() error) error {
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

// OnceValue is similar to Once, except it returns a value.
type OnceValue struct {
	m     sync.Mutex
	done  uint32
	value interface{}
}

// Do runs the specified function only once, but all callers gets the same
// result from that single execution.
func (o *OnceValue) Do(f func() interface{}) interface{} {
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

// OnceValueError is similar to Once, except it return a (value, error) tuple
type OnceValueError struct {
	m     sync.Mutex
	done  uint32
	value interface{}
	err   error
}

// Do runs the specified function only once, but all callers gets the same
// result from that single execution.
func (o *OnceValueError) Do(f func() (interface{}, error)) (interface{}, error) {
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
