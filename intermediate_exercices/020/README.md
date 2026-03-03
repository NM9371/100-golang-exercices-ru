# Упражнение: Структуры

## Что такое структуры?

В Go нет ключевого слова `class` для определения объектов. В Go концепция объединения полей разных типов данных в одну переменную называется структурой (struct).

Структура — это типизированная коллекция полей.
Она может хранить данные любого типа, в том числе другие структуры.
Пример структуры:

```go
type Person struct {
  name string
  age int
  job string
  salary int
}
```

Объект `Person` содержит 4 поля: `name`, `age`, `job` и `salary`.
Поля могут быть разных типов данных.

## Доступ к полям структуры

Для доступа к любому полю структуры используйте оператор точки (.) между именем переменной-структуры и именем поля:

```go
func main() {
  var pers1 Person

  // Задаём данные Pers1
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 1000

  // Читаем и выводим данные Pers1
  fmt.Println("Name: ", pers1.name) // Output: 'Name: Hege'
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)
}
```

## Анонимные структуры

Структура без имени называется анонимной. Их часто можно встретить в [табличных тестах](https://go.dev/wiki/TableDrivenTests).

Анонимные структуры определяются без имени и поэтому не могут быть использованы в других местах кода.
Чтобы создать анонимную структуру, достаточно сразу создать её экземпляр, добавив вторую пару фигурных скобок после объявления типа:

```go
newCar := struct {
    make    string
    model   string
    mileage int
}{
    make:    "Ford",
    model:   "Taurus",
    mileage: 200000,
}
```

Упражнение: Создайте структуру `Hotel` с тремя полями: `numRooms`, `streetName` и `hasPool`

- numRooms int32
- streetName string
- hasPool bool

Затем присвойте значение каждому из этих полей.

Ссылки:
[https://go.dev/tour/moretypes/2](https://go.dev/tour/moretypes/2)
[https://www.w3schools.com/go/go_struct.php](https://www.w3schools.com/go/go_struct.php)
[https://gobyexample.com/structs](https://gobyexample.com/structs)

```go
package main

import "fmt"

type Hotel struct {
  // Ваш код здесь

}

func main () {
  var myHotel Hotel
  // Ваш код здесь
}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

type Hotel struct {
  // Ваш код здесь
  numRooms int32
  streetName string
  hasPool bool
}

func main () {
  var myHotel Hotel
  // Ваш код здесь
  myHotel.numRooms = 30
  myHotel.streetName = "Thaerstrasse"
  myHotel.hasPool    = true
  fmt.Printf("My hotel in %v has %d rooms and it's %t that has a Pool", myHotel.streetName, myHotel.numRooms, myHotel.hasPool)
}
```

</details>
