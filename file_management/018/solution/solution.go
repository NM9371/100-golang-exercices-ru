// Упражнение: Запишите список из 5 стран в файл
// Подсказка: используйте пакет "os"

package main

import "os"

func main() {
	// Ваш код здесь
	file, err := os.Create("contries.txt")

	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("Germany\nFrance\nUSA\nSpain\nUK\n")
}
