// Упражнение: Случайные числа
// Сгенерируйте случайное число из диапазона [-50, +50]
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(-50, +51)
	fmt.Printf("Random number: %d\n", randomNum)
}
