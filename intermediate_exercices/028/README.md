# Упражнение: Вложенные функции

Функция — это блок операторов, который можно многократно использовать в программе.

Функции очень удобны и широко применяются в программировании — они облегчают чтение, понимание и сопровождение кода. Они также способствуют повторному использованию кода [принцип DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)

## Анатомия функции

```go
func functionName(param1 type1, param2 type2) returnType { // Это называется сигнатурой функции
    // Внутри фигурных скобок находится тело функции с логикой.
}
```

Пример:

```go
func subtract(x,y int) int { // сигнатура функции
    return x - y // тело функции, возвращает вызывающей стороне результат вычитания x - y.
}

func main() {
    result := subtract(5,4) // вызов функции, её возвращаемое значение присваивается переменной `result`
    fmt.Println(result) // результат будет 1
}
```

Создайте функцию, складывающую 2 числа, а затем вызовите её из другой функции, которая суммирует ещё одно число.

```go
package main

import "fmt"

func sum() int {

}

func secondsum() int {

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

func sum(x int, y int) int {
    return x + y
}

func secondsum(x, y, z int) int {
    return sum(x,y) + z
}

func main () {
    // Ваш код здесь
    var result int
    result = secondsum(2,3,5)

    fmt.Println(result)
}
```

</details>
