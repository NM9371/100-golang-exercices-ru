// Упражнение: Каналы и таймауты

// Создайте горутину, которая использует time.Sleep на 10 секунд, затем добавляет строку "10 seconds passed" в канал (типа string)
// в главной программе внутри блока select имейте 2 случая:
// 1- Сообщение из канала
// 2- Таймаут с оператором "time.After(3 * time.Second)". После таймаута выведите "Timeout!!!!"

package main

import "fmt"
import "time"

func timeout(c chan string) {
	for {

	}
}

func main() {
	var c1 chan string = make(chan string)

	go timeout(c1)
	for {
		select {}
	}
	fmt.Println("Goroutines finished.") // Это сообщение не должно появиться, т.к. горутины работают вечно!
}
