# gosafe

[![Go Reference](https://pkg.go.dev/badge/github.com/janitorjeff/gosafe.svg)](https://pkg.go.dev/github.com/janitorjeff/gosafe)

Store values in a thread-safe way. Supports arbitrary values, slices and maps.
Uses generics which means that go 1.18+ is required.

## Example

```go
package main

import (
	"github.com/kvlach/gosafe"
)

func main() {
	m := gosafe.Map[int, string]{} // equivalent to map[int]string
	m.Set(123, "abc")              // m[123] = "abc"
	n, ok := m.Get(123)            // n, ok := m[123]
	// ...
}
```

Alongside the thread-safe methods unsafe ones are provided which allow for easy
extension.

```go
package main

import (
	"github.com/kvlach/gosafe"
)

type CustomSlice struct {
	gosafe.Slice[string]
}

func NewSlice() CustomSlice {
	return CustomSlice{gosafe.Slice[string]{}}
}

func (s *CustomSlice) ChangeElemsWithLengthGreaterThanThree() {
	// make it thread safe
	s.Lock()
	defer s.Unlock()

	// use unsafe methods to avoid mutex deadlock
	for i := 0; i < s.LenUnsafe(); i++ {
		if len(s.GetUnsafe(i)) > 3 {
			s.SetUnsafe(i, "changed")
		}
	}
}

func main() {
	s := NewSlice()
	s.Append("abc", "def", "long-string", "xyz")
	s.ChangeElemsWithLengthGreaterThanThree()
	print(s.Get(0), "\n") // abc
	print(s.Get(2))       // changed
}
```
