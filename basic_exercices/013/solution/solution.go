// Упражнение: Условные операторы
// Проверьте, находится ли число в диапазоне от 20 до 30
// Если число меньше 20, выведите: too cold
// Если число в диапазоне, выведите: perfect
// Если число больше 30, выведите: so hot

// Используйте if и else if!

package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	if number < 20 {
		fmt.Println("Too cold!")
	} else if number > 20 && number < 30 {
		fmt.Println("perfect")
	} else {
		fmt.Println("so hot!")
	}

}
