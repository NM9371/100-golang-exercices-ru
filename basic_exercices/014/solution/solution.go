// Упражнение: Пользовательский ввод
// Используя только пакет fmt, запросите у пользователя имя, а затем фамилию
// После ввода данных пользователем выведите их

// Подсказка: https://pkg.go.dev/fmt#hdr-Scanning

package main

import "fmt"

func main() {
	// Здесь пишите ваш код
	var name, surname string
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("Please enter your surname")
	fmt.Scanln(&surname)

	fmt.Printf("Your name is: " + name + " " + surname)
}
