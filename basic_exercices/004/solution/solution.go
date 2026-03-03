package main

import "fmt"

func main() {
	// Создаём новую переменную с именем helloWorld
	var helloWorld, itsMeMario string
	helloWorld = "Hello World!"
	itsMeMario = "It's a me, Mario"
	// Выводим переменную
	fmt.Println(helloWorld + " " + itsMeMario)
}

// Чтобы запустить программу:
// - go run solution.go
