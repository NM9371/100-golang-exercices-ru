// Упражнение: Чтение файла
// Подсказка: используйте пакет "os"

// Внимание: запускайте этот код там, где находится читаемый файл, или укажите полный путь!
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("/tmp/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) // or os.Stdout.Write(data)
}
