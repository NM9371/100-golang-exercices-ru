// Упражнение: Изучите ФОРМАТНЫЕ СТРОКИ! Выведите несколько глаголов с соответствующими им буквами
// Форматные строки / глаголы предваряются символом '%'.
// Пример: fmt.Printf("Hello, my name is %s", name) // подставит вместо %s содержимое переменной с именем "name"

// Подсказка: Обратитесь к документации: https://pkg.go.dev/fmt#hdr-Printing

package main

import "fmt"

func main() {
	// Здесь пишите ваш код
	var name string
	var age int64
	var legal bool
	var weight float32

	name = "Anna"
	age = 20
	legal = true
	weight = 70.12

	fmt.Printf("My name is %s, I am %d years old and it's %t that I can drive a car, my pet weights %f kilograms", name, age, legal, weight)
}
