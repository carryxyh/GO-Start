package main

import (
	"fmt"
	"testing"
)

type Animal interface {
	Loud()
}

type Cat struct {
	miao string
}

// go:noinline
func (c Cat) Loud() {
	fmt.Println(c.miao)
}

// go:noinline
func BenchmarkTestInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c Animal = Cat{
			"miao",
		}
		c.Loud()
	}
}
