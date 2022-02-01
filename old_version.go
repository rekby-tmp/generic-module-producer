//go:build !go1.18

package generic_module_producer

import "fmt"

func Print(v interface{}) {
	fmt.Printf("old %v\n", v)
}
