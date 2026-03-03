# Упражнение: Функции III

Вариадические функции могут вызываться с произвольным числом аргументов.
В Go для создания вариадических функций используется многоточие `...` (Ellipsis).

Согласно [спецификации Go](https://go.dev/ref/spec#Function_types):
> Последний входной параметр в сигнатуре функции может иметь тип с префиксом .... Функция с таким параметром называется вариадической и может быть вызвана с нулём или более аргументами для этого параметра.

Функция, объявленная с многоточием, может принимать любое количество аргументов:

```go
package main

import "fmt"

func main() {
  // 0 аргументов
    sayHello()
  // 1 аргумент
    sayHello("Rahul")
  // 4 аргумента
    sayHello("Mohit", "Rahul", "Rohit", "Johny")
}

// Использование многоточия (Ellipsis)
func sayHello(names ...string) {
    for _, n := range names {
        fmt.Printf("Hello %s\n", n)
    }
}
```

Создайте вариадическую функцию, которая суммирует все переданные ей числа.

```go
package main

import "fmt"

// Дополните сигнатуру функции
func sum() int {
  total := 0
    // Ваш код здесь
}

func main () {
  // Ваш код здесь

}

```

<details>
<summary> Решение: </summary>

```golang
package main

import "fmt"

func sum(numbers ...int) int {
  total := 0
  for _,num := range numbers{
    total += num
  }
  fmt.Println(total)
  return total
}

func main () {
  // Ваш код здесь
  sum(2,3,4,5,6,7)
}
```

</details>
