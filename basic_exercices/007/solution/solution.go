package main

import "fmt"

func main() {
	// Создаём новую переменную с именем helloWorld
	var helloWorld string
	helloWorld = "Hello World!"
	// Выводим первую букву
	fmt.Println(helloWorld[0])
}

// Чтобы запустить программу:
// - go run solution.go
