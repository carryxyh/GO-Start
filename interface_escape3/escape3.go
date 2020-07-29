package main

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c Cat) Quack() {
	println(c.Name + " meow")
}

func main() {
	var c *Cat = &Cat{Name: "grooming"}
	Qua(c)
}

//go:noinline
func Qua(cat *Cat) {
	cat.Quack()
}
