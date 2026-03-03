# Упражнение: Форматные глаголы

Go предоставляет удобный способ форматировать вывод с помощью форматных глаголов.
Вы указываете программе, как хотите представить переменную в соответствии с её типом.

Форматный глагол действует как заполнитель для переменной, которую вы в него передаёте. Вот два примера:

Глагол `%v` выводит значение в формате по умолчанию

```go
fmt.Printf("my string variable contains %v", dog_name)
```

Выведет "my string variable contains Scooby"

Глагол `%T` выводит тип переменной

```go
fmt.Printf("my string variable is of type %T", dog_name)
```

Выведет "my string variable is of type string"

Все форматные глаголы можно найти в документации: https://pkg.go.dev/fmt#hdr-Printing

Упражнение: Выведите несколько форматных глаголов с соответствующими им буквами

- Форматные глаголы предваряются символом '%'.

```go
package main

import "fmt"

func main () {
  var name string
  var age int64
  var legal bool
  var weight float32

  name = "Anna"
  age  = 29
  legal = false
  weight = 70.12

  // Здесь пишите ваш код
  fmt.Printf("My name is __, I am __ years old and it's __ that I can drive a car, my pet weights __ kilograms", , , , )
}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

func main () {
  // Здесь пишите ваш код
  var name string
  var age	int64
  var legal bool
  var weight float32

  name = "Anna"
  age  = 20
  legal = true
  weight = 70.12

  fmt.Printf("My name is %s, I am %d years old and it's %t that I can drive a car, my pet weights %f kilograms",name, age, legal, weight)
}
```

</details>
