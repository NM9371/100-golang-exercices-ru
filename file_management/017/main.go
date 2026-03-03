// Упражнение: Чтение файла

// Внимание: запускайте этот код там, где находится читаемый файл, или укажите полный путь!
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
