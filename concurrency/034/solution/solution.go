// Упражнение: Каналы (буферизованные)

// Каналы, принимающие одно сообщение (<-), синхронны. Но Go также может использовать асинхронные или буферизованные каналы
// Создайте буферизованный строковый канал с ёмкостью 4
// Отправьте в него напрямую 4 разные строки.
// Используйте функцию pop_message 4 раза, чтобы извлечь элементы из канала и посмотреть, как это работает

package main

import "fmt"
import "time"

func pop_message(c chan string) {
	msg := <-c
	fmt.Println(msg)
}
func main() {
	// Ваш код здесь
	var c chan string = make(chan string, 4)
	c <- "My"
	c <- "Name"
	c <- "is"
	c <- "Enin"

	pop_message(c)
	pop_message(c)
	pop_message(c)
	pop_message(c)

	// эта пауза нужна, чтобы не завершить программу раньше времени :)
	time.Sleep(1 * time.Second)
}
