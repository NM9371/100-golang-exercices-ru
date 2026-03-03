// Упражнение: Функции

// Создайте вариадическую функцию, которая суммирует переданные ей числа
// Вариадические функции могут вызываться с любым числом аргументов.

package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println(total)
	return total
}

func main() {
	// Ваш код здесь
	sum(2, 3, 4, 5, 6, 7)
}
