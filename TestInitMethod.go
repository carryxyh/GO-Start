package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

type DBconfig struct {
	Alias      string
	DriverName string
	DataSource string
	Conns      []int
}

func main() {
	var c *DBconfig
	//fmt.Println(c.Alias)  <- 这里报错
	fmt.Println(&c)

	// cc := DBconfig{}

	fmt.Println(new(DBconfig))
	fmt.Println("main...")

	t := DBconfig{}
	fmt.Println(&t)
}
