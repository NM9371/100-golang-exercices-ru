# Упражнение: Итоговое упражнение

Это последнее упражнение блока basic_exercises. Вы применёте на практике всё, что узнали.

В этом упражнении мы проверим, ввёл ли пользователь число и является ли это число чётным.

Чётное число — это число, делящееся на 2 без остатка.
Для получения остатка от деления используется [оператор остатка/модуля](https://en.wikipedia.org/wiki/Modulo) `%`.

Для простоты мы не будем проверять, является ли ввод числом — предполагается, что пользователь всегда вводит число.

- Определите переменную типа `int32` с именем `number`
- Создайте функцию для проверки чётности числа; функция будет называться `isEven`
- Функция выводит в стандартный вывод, является ли число чётным или нечётным.
- Функция возвращает значение типа `bool`.

```go
package main

import "fmt"

func main () {
  // Здесь пишите ваш код
  fmt.Printf("...")
}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

func main () {
  var number int32
  fmt.Println("Enter a number: ")
  fmt.Scanln(&number)

  isEven(number)
}

func isEven(number int32) bool {
  if (number % 2 == 0) {
    fmt.Println("is even")
    return true
  }  else {
    fmt.Println("is odd")
    return false
  }
}
```

</details>
