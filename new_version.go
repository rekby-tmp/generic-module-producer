//go:build go1.18
// +build go1.18

package generic_module_producer

import "fmt"

type Ti interface {
	~int | ~int8
}

func Print[T Ti](v T) {
	fmt.Printf("new %v\n", v)
}
