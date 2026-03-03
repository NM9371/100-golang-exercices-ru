// Упражнение: Пользовательский ввод
// Получите число из консоли и проверьте, является ли оно нечётным

// Можно создать отдельную функцию или разместить всё в функции main :)

package main

import "fmt"

func main() {
	var number int32
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	iseven(number)

	/*
		Возможное решение без дополнительной функции:
		if (number % 2 == 0) {
			fmt.Println("It's even")
		} else {
			fmt.Println("It's odd")
		}*/
}

func iseven(number int32) bool {
	if number%2 == 0 {
		fmt.Println("is even")
		return true
	} else {
		fmt.Println("is odd")
		return false
	}
}
