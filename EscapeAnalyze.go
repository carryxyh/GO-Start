package main

type Cursor struct {
	X int
}

func main() {
	f()
}

func f() *Cursor {
	var c Cursor
	c.X = 500
	// noinline()
	return &c
}
