// Упражнение: Несколько разных каналов

// Используем оператор 'select'
// Создайте три разных канала "c1,c2,c3" для горутины с именем "name"
// Горутина будет иметь 2 аргумента: сначала канал (тип string), затем имя (тип string)
// Внутри горутины (функции) будет:
// 1- Бесконечный цикл:
// 1.1- Второй аргумент "name" будет записан в канал
// 1.2- Подождать 2 секунды

// В главной программе
// Создайте три разных канала "c1,c2,c3" типа string
// вызовите горутину 'name' 3 раза с разным именем каждый раз
// создайте бесконечный цикл, содержащий оператор select
// Внутри select поместите 3 ключевых слова 'case'. В каждом случае выведите имя для каждого разного канала.

package main

import "fmt"
import "time"

func name(c chan string, name string) {
	for {
		c <- name
		time.Sleep(2 * time.Second)
	}
}

func main() {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)
	var c3 chan string = make(chan string)

	go name(c1, "John")
	go name(c2, "Mary")
	go name(c3, "Fritz")

	for {
		select {
		case name1 := <-c1:
			fmt.Println(name1)
		case name2 := <-c2:
			fmt.Println(name2)
		case name3 := <-c3:
			fmt.Println(name3)
		}
	}
	fmt.Println("Goroutines finished.") // Это сообщение не должно появиться, т.к. горутины работают вечно!
}
