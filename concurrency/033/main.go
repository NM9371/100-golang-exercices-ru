// Упражнение: Каналы

// Важно понять эту концепцию, сначала изучите самостоятельно! :)
// Есть две конкурентные горутины (f1 и f2)
// Отправьте сообщение "Hello from f1" из функции f1
// Получите сообщение из f1 в функции f2 и выведите "I am f2 and ..." + сообщение от f1
// Сделайте это с помощью канала

package main

import "fmt"
import "time"

func f1(c chan string) {

}

func f2(c chan string) {

}
func main() {

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(1 * time.Second)
}
