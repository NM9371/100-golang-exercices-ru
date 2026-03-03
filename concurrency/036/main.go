// Упражнение: Синхронизация каналов

// Есть две горутины
// Вызовите первую горутину (с именем 'task') и запустите вторую горутину 'task2' ТОЛЬКО после завершения первой

package main

import "fmt"
import "time"

func task(done chan bool) {
	fmt.Print("running Task1 goroutine...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func task2() {
	fmt.Println("Task2 goroutine...")
}

func main() {
	var done chan bool = make(chan bool, 1)
	fmt.Println("I am running in the main thread concurrently")
	// Ваш код здесь

}
