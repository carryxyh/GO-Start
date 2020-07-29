package main

import (
	"testing"
)

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c *Cat) Quack() *Cat {
	println(c.Name + " meow")
	return c
}

//go:noinline
func Qua(cat *Cat) {
	cat.Quack()
}

// go:noinline
func BenchmarkTestInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c Duck = &Cat{
			"tom",
		}
		Qua(c)
	}
}
