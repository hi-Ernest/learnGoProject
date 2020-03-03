package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool
var i int
//fmt.Printf(i, c, python, java);
//0, false, false, false

var t, j int = 1, 3


func main() {
	fmt.Println("hello world");
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(add(2, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//仅仅在函数内 -> 赋值
	i:= 32
	fmt.Print(i)

}

func add(x int, y int) int {
	return x + y
}

func remove(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


