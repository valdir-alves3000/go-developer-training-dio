package main

import "fmt"

func main() {
	x := Sum(1, 2, 3)
	y := Multiply(10, 10)
	w := Subtract(5, 10)
	z := Divide(20, 4)
	fmt.Println(x, y, w, z)
}

func Sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func Subtract(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	total := nums[0]
	for _, v := range nums[1:] {
		total -= v
	}
	return total
}

func Multiply(nums ...int) int {
	total := 1
	for _, v := range nums {
		total *= v
	}
	return total
}

func Divide(a, b int) int {
    if b == 0 {
        panic("Divisão por zero não é permitida")
    }
	return a / b
}