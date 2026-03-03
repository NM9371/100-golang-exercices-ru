// Упражнение: Функции

// Создайте функцию, возвращающую 2 целочисленных значения
// Функция принимает 2 аргумента (int)
// Первое возвращаемое значение — сумма аргументов, второе — их разность

package main

import "fmt"

func operation(x, y int) (int, int) {
	var sum, substraction int
	sum = x + y
	substraction = x - y
	return sum, substraction
}

func main() {
	// Ваш код здесь
	sum, subs := operation(10, 5)
	fmt.Println(sum, subs)
}
