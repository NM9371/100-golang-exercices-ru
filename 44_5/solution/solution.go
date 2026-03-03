// Упражнение: Финальное упражнение: Конкурентность

// Создайте 2 горутины: одну с именем number() и другую с именем letter(), без аргументов и без возвращаемых значений.
// Функция number() будет печатать числа от 0 до 9, ожидая 200 миллисекунд между каждой итерацией.
// Функция letter() будет печатать первые 10 букв алфавита, ожидая 200 миллисекунд между каждой итерацией.
// В главной программе запустите горутины и выведите сообщение, наблюдая за поведением программы.

package main

import (
	"fmt"
	"time"
)

func number() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func letter() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%c ", ('a' + i))
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go number()
	go letter()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
}
