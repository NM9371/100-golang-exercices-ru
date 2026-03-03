// Упражнение: СТРУКТУРЫ
// Создайте структуру Hotel со следующими полями:
// numRooms int32
// streetName string
// hasPool bool

// Затем присвойте значение каждому из этих полей

package main

import "fmt"

type Hotel struct {
	// Your code goes here
	numRooms   int32
	streetName string
	hasPool    bool
}

func main() {
	var myHotel Hotel
	// Ваш код здесь
	myHotel.numRooms = 30
	myHotel.streetName = "Thaerstrasse"
	myHotel.hasPool = true
	fmt.Printf("My hotel in %v has %d rooms and it's %t that has a Pool", myHotel.streetName, myHotel.numRooms, myHotel.hasPool)
}
