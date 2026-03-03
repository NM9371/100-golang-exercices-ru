// Упражнение: Указатели

// Объявите переменную-указатель типа int32
// Получите адрес переменной "x"
// Сохраните адрес переменной "x" в переменную-указатель

package main

import "fmt"

func main() {
	var x int32 = 5
	// Ваш код здесь
	var pointerX *int32 = &x

	fmt.Println("Value of x: %d", x)
	fmt.Println("Memory address of x: %d", &x)
	fmt.Println("Pointer value: %d", pointerX)
}
