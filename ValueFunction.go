/* 函数作为值 */
package main

import "fmt"
import "math"

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}
