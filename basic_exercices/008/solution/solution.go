// Упражнение: Массивы
// Создайте массив из 10 значений типа "int8" и при инициализации заполните его значениями от 0 до 9

package main

import "fmt"

func main() {
	var arr = [5]string{"thomas", "phillip"}

	// Выводим массив
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
