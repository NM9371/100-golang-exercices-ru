// Упражнение: Каналы — range (небуферизованный)

// В этом упражнении мы будем использовать ключевое слово range для итерации по НЕбуферизованному (синхронному) каналу.
// Создайте небуферизованный канал (тип int)
// Создайте горутину, которая:
// Имеет бесконечный цикл -> выводит текущую секунду и ждёт одну секунду
// Используйте итератор range в функции main, чтобы видеть каждую секунду

package main

import "fmt"
import "time"

func second(c chan int) {
	for {
		c <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func main() {
	var c chan int = make(chan int)

	go second(c)

	for element := range c {
		fmt.Println(element)
	}
}
