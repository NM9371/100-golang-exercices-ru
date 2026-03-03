// Упражнение: Каналы

// Важно понять эту концепцию, сначала изучите самостоятельно! :)
// Есть две конкурентные горутины (f1 и f2)
// Отправьте сообщение "Hello from f1" из функции f1

package main

import "fmt"
import "time"

func f1(c chan string) {
	c <- "Hello from f1"
}

func f2(c chan string) {
	msg := <-c
	fmt.Println("I am f2 and...", msg)
}
func main() {
	// Ваш код здесь
	var c chan string = make(chan string)
	go f1(c)
	go f2(c)

	// эта пауза нужна, чтобы не завершить программу раньше времени :)
	time.Sleep(1 * time.Second)
}
