# Упражнение: Условные операторы

Условные операторы — ещё одна фундаментальная тема в программировании и строительные блоки алгебраической и пропозициональной логики.
Условные операторы используются для выполнения различных действий в зависимости от условий.

Условие может быть `true` или `false`, и в зависимости от результата выполняется уникальная ветка выполнения, как показано на этом изображении:

![if-then chart](if-then.png)

Go поддерживает стандартные операторы сравнения из математики:
Меньше `<`
Меньше или равно `<=`
Больше `>`
Больше или равно `>=`
Равно `==`
Не равно `!=`

Дополнительно Go поддерживает стандартные логические операторы:
Логическое И `&&`
Логическое ИЛИ `||`
Логическое НЕ `!`

Вот пример кода, демонстрирующий работу условных операторов:

```go
var speed=50

// Одно условие — ничего не происходит, если оно не выполнено.
if (speed > 60 ){
  fmt.Printf("Only fast animals will see this message")
}
// Вывод:

// Одно условие — если не выполнено, выполняется блок `else`.
if (speed < 50){
  fmt.Printf("Going slow")
} else {
  fmt.Printf("You are a fast animal!")
}
// Вывод: You are a fast animal!

// Несколько условий
if (speed < 20){
  fmt.Printf("Going slow")
} else if ( speed >= 20 && speed < 30 ) {
  fmt.Printf("You are fast")
} else if ( speed >= 30 && speed < 40 ) {
  fmt.Printf("You are super fast")
} else if ( speed >= 40 && speed <= 80 ) {
  fmt.Printf("You are mega super fast")
} else {
  fmt.Printf("You are the fastest of 'em all!")
}
// Вывод: You are mega super fast
```

Операторы и их комбинации можно использовать для создания условий при различных решениях, как мы видели с `( speed >= 40 && speed <= 80 )`.
Операторы также имеют приоритет, который необходимо учитывать. Его можно найти в [спецификации go](https://go.dev/ref/spec#Operator_precedence); изменить приоритет можно, добавив скобки вокруг операций, которые нужно выполнить вместе.

```go
42 + 2 * 8  // 42 + (2 * 8) = 58
(42 + 2) * 8 // (42 + 2) * 8 = 352
```

В приведённом выше коде приоритет операторов изменён с помощью скобок

Упражнение:

- Проверьте, находится ли число в диапазоне от 20 до 30
- Если число меньше 20, выведите: too cold
- Если число в диапазоне, выведите: perfect
- Если число больше 30, выведите: so hot

```go
// Используйте if и else if!
package main

import "fmt"

func main () {
  var number int
  fmt.Scanln(&number)

  if (number < 20){
    fmt.Println("Too cold!")
  } else if (number > 20 && number < 30) {
    fmt.Println("perfect")
  } else {
    fmt.Println("so hot!")
  }
}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

func main () {
  var number int
  fmt.Scanln(&number)

  if (number < 20){
    fmt.Println("Too cold!")
  } else if (number > 20 && number < 30) {
    fmt.Println("perfect")
  } else {
    fmt.Println("so hot!")
  }

}
```

</details>
