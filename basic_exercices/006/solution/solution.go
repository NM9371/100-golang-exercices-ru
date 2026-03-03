// Упражнение: Массивы
// Создайте массив из 10 значений типа "int8" и при инициализации заполните его значениями от 0 до 9

package main

import "fmt"

func main() {
	var arr = new([10]int)
	/*
		Простое, но ручное решение:
		arr[0] = 0
		arr[1] = 1
		arr[2] = 2
		arr[3] = 3
		arr[4] = 4
		arr[5] = 5
		arr[6] = 6
		arr[7] = 7
		arr[8] = 8
		arr[9] = 9
	*/

	// Ещё одно решение :)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
		fmt.Println(arr[i])
	}
}
