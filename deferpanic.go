package main

import (
	"fmt"
)

// func main() {
// 	defer recoverSth()
// 	defer fmt.Println("in main")
// 	defer func() {
// 		// defer func() {
// 		// 	panic("panic again and again")
// 		// }()
// 		// panic("panic again")
// 		defer recoverSth()
// 		panic("again?")
// 	}()

// 	panic("panic once")
// }

// func recoverSth() {
// 	if err := recover(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func main() {
// 	defer fmt.Println("in main")
// 	defer func() {
// 		defer func() {
// 			panic("panic again and again")
// 		}()
// 		panic("panic again")
// 	}()

// 	panic("panic once")
// }

// func main() {
// 	fmt.Println("sth")
// 	defer recoverSth()
// 	err, boo := t(1)
// 	if err != nil && boo {
// 		panic(err)
// 	}
// }

// func t(i int) (error, bool) {
// 	if i == 1 {
// 		return errors.New("bad input"), true
// 	} else {
// 		return nil, false
// 	}
// }

// func recoverSth() {
// 	if err := recover(); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
