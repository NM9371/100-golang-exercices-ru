// Упражнение: Функции

// Создайте рекурсивную функцию, возвращающую последовательность Фибоначчи до n-го числа

package main

import "fmt"

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	// Ваш код здесь
	fmt.Println(fibonacci(9))
}
