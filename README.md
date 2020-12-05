# hsson/once
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hsson/once)](https://pkg.go.dev/github.com/hsson/once) [![GoReportCard](https://goreportcard.com/badge/github.com/hsson/once)](https://goreportcard.com/report/github.com/hsson/once)

A re-implementation and drop-in replacement of the standard `sync.Once`, with added support for return values! This package exports three additional `Once`-like primitives, in addition to the standard `once.Once`:

`once.Error` returns an error value
```go
Do(f func() error) error
```
`once.Value` returns a value
```go
Do(f func() interface{}) interface{}
```
`once.ValueError` returns a (value, error) tuple
```go
Do(f func() (interface{}, error)) (interface{}, error)
```

These three primitives have the behavior that, like with the standard `Once`, the function passed is ever only executed once. However, they also return the value returned by that one execution to all subsequent callers.