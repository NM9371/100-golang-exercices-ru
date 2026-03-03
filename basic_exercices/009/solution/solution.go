// Упражнение: Циклы for
// Пройдитесь по массиву и добавьте два (+2) к каждому значению

package main

import "fmt"

func main() {

	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + 2
		fmt.Println(arr[i])
	}
}
