// Упражнение: Функции

// Создайте функцию, складывающую 2 числа, затем вызовите её из другой функции, которая складывает ещё одно число.

package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func secondsum(x, y, z int) int {
	return sum(x, y) + z
}

func main() {
	// Ваш код здесь
	var result int
	result = secondsum(2, 3, 5)

	fmt.Println(result)
}
