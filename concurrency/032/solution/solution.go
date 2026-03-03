// Упражнение: Горутины

// Изучите самостоятельно, что такое горутина.
// Создайте горутину с именем `start()` и вызовите её внутри блока main()
// Обратите внимание, как горутины не ждут завершения выполнения и продолжают со следующих инструкций

package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func start() {
	fmt.Println("In Goroutine")
}
