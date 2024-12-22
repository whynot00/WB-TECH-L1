package main

import (
	"fmt"
	"math/big"
)

func bigDiv(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func bigSub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func bigMul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func bigAdd(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func main() {
	// Создаем большие целые числа a и b
	a := big.NewInt(1 << 21) // a = 2^21
	b := big.NewInt(1 << 22) // b = 2^22

	// Умножение
	fmt.Printf("Результат умножения a * b: %v\n", bigMul(a, b))

	// Деление
	fmt.Printf("Результат деления a / b: %v\n", bigDiv(a, b))

	// Сложение
	fmt.Printf("Результат сложения a + b: %v\n", bigAdd(a, b))

	// Вычитание
	fmt.Printf("Результат вычитания a - b: %v\n", bigSub(a, b))
}
