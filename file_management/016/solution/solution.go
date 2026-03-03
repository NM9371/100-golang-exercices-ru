// Упражнение: Проверьте существование файла
// Подсказка: используйте пакет "os"

package main

import "fmt"
import "os"

func main() {
	// Ваш код здесь
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}
