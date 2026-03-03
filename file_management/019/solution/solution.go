// Упражнение: Переименуйте файл с name1 на name2
// Подсказка: используйте пакет "os"

package main

import "os"

func main() {
	var src, dest string
	// Ваш код здесь

	src = "name1.txt"
	dest = "name2.txt"

	os.Rename(src, dest)
}
