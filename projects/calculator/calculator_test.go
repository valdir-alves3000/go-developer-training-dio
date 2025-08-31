package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"soma de números positivos", []int{1, 2, 3}, 6},
		{"lista vazia", []int{}, 0},
		{"números negativos", []int{-1, -2, -3}, -6},
	}

	for _, trun := range tests {
		t.Run(trun.name, func(t *testing.T) {
			got := Sum(trun.nums...)
			if got != trun.want {
				t.Errorf("Sum(%v) = %d; want %d", trun.nums, got, trun.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"subtração simples", []int{10, 5}, 5},
		{"lista vazia", []int{}, 0},
		{"subtração de múltiplos", []int{20, 5, 3}, 12},
	}

	for _, trun := range tests {
		t.Run(trun.name, func(t *testing.T) {
			got := Subtract(trun.nums...)
			if got != trun.want {
				t.Errorf("Subtract(%v) = %d; want %d", trun.nums, got, trun.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"multiplicação simples com resultado 6", []int{2, 3}, 6},
		{"multiplicação simples com resultado 15", []int{5, 3}, 15},
		{"lista vazia", []int{}, 1},
		{"com zero", []int{2, 0, 5}, 0},
		{"com zero", []int{2, 10, 0}, 0},
		{"números negativos com resultado -6", []int{-2, 3}, -6},
		{"números negativos com resultado -8", []int{2, -4}, -8},
	}

	for _, trun := range tests {
		t.Run(trun.name, func(t *testing.T) {
			got := Multiply(trun.nums...)
			if got != trun.want {
				t.Errorf("Multiply(%v) = %d; want %d", trun.nums, got, trun.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"divisão simples com resultado 5", 20, 4, 5},
		{"divisão simples com resultado 3", 21, 7, 3},
		{"divisão com negativo", -20, 4, -5},
		{"divisão com negativo", 20, -4, -5},
	}

	for _, trun := range tests {
		t.Run(trun.name, func(t *testing.T) {
			got := Divide(trun.a, trun.b)
			if got != trun.want {
				t.Errorf("Divide(%d, %d) = %d; want %d", trun.a, trun.b, got, trun.want)
			}
		})
	}
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r != "Divisão por zero não é permitida" {
			t.Errorf("esperava panic com mensagem 'Divisão por zero não é permitida', mas recebeu: %v", r)
		}
	}()
	Divide(10, 0)
}