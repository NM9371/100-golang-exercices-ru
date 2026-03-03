// Упражнение: Синхронизация каналов

// У нас есть горутина task, изучите, как работает синхронизация через канал, и заставьте программу ждать асинхронную функцию (сделайте её синхронной ;) )
// Это упражнение простое, вам нужно добавить всего одну строку

package main

import "fmt"
import "time"

func task(done chan bool) {
	fmt.Print("running...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	var done chan bool = make(chan bool)
	go task(done)

	// Ваш код здесь

}
