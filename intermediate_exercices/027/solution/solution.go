// Упражнение: СРЕЗЫ

// Создайте срез из среза, выбрав средние 3 элемента

package main

import "fmt"

func main() {
	var slice = []int32{0, 1, 2, 3, 4}
	// Ваш код здесь
	// Помните, что индекс среза начинается с 0!
	new_slice := slice[1:3]
	fmt.Printf("substring:%v, type: %T", new_slice, new_slice)
}
