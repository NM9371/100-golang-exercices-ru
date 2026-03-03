// Упражнение: Каналы — закрытие

// Создайте строковый канал "c" (сделайте его буферизованным)
// Добавьте в этот канал 2 разные строки напрямую.
// Закройте канал оператором close() и прочитайте значение из канала. Можно ли его прочитать?

package main

import "fmt"

func main() {
	var c chan string = make(chan string, 2)

	c <- "Hello"
	c <- "how are you?"

	fmt.Println(<-c)
	close(c)
	fmt.Println(<-c)
}
