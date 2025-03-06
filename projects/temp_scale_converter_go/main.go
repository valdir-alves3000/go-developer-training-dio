package main

import "fmt"

const (
	colorRed   = "\033[31m"
	colorBlue  = "\033[34m"
	colorReset = "\033[0m"

	kelvin   = 273
	boilingK = 373
	fusionK  = 273
)

func main() {
	boilingC := boilingK - kelvin
	fusionC := fusionK - kelvin

	boilingF := (boilingC * 9 / 5) + 32
	fusionF := (fusionC * 9 / 5) + 32

	fmt.Println(colorRed + "Ponto de Ebulição da Água:" + colorReset)
	fmt.Printf("Kelvin: %v Kº\n", boilingK)
	fmt.Printf("Celsius: %v Cº\n", boilingC)
	fmt.Printf("Fahrenheit: %v Fº\n\n", boilingF)

	fmt.Println(colorBlue + "Ponto de Fusão da Água:" + colorReset)
	fmt.Printf("Kelvin: %v Kº\n", fusionK)
	fmt.Printf("Celsius: %v Cº\n", fusionC)
	fmt.Printf("Fahrenheit: %v Fº\n", fusionF)
}