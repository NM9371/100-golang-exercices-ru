// Упражнение: СРЕЗЫ

// Создайте срез из первых 5 чисел из списка из 10 чисел

package main

import "fmt"

func main() {
	var myset = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Ваш код здесь
	var slice = myset[0:5]
	fmt.Println("Slice :", slice)
}
