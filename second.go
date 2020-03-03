package main

import (
	"fmt"
	"math"
)

const Pi  = 3.14

func main() {
	const World = "世界"
	const Truth = true
	fmt.Println("Hello", World)

	fmt.Println("happy", Pi, World)

	fmt.Println("Go rules ?", Truth)

	sum := 0
	for i:=0 ; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	tem := 1
	for tem < 1000 {
		tem += tem
	}
	fmt.Println(tem)

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))


}

func pow(x, n, lim float64) (float64)  {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}
