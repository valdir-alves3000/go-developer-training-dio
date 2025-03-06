package main

import "fmt"

func main() {
	output := ""

	for i := 1; i <= 100; i++ {
		if isMultiple3(i) && isMultiple5(i) {
			output = "PinPan"
		} else if isMultiple3(i) {
			output = "Pin"
		} else if isMultiple5(i) {
			output = "Pan"
		} else {
			output = fmt.Sprintf("%d", i)
		}
		fmt.Println(output)
	}
}

func isMultiple3(num int) bool {
	return num%3 == 0
}

func isMultiple5(num int) bool {
	return num%5 == 0
}
