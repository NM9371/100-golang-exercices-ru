# Упражнение: Функции II

Функции также могут возвращать несколько значений, как в этом примере:

```go
func convertCaps (name string) (string, string) {
  // преобразуем имя в верхний регистр и возвращаем оригинал и преобразованную строку
  transformed := strings.ToUpper(name)
  return name, transformed
}

name, upperName := convertCaps("john")
fmt.Printf("original: %v, transformed: %v", name, upperName)
// Output: original: john, transformed: JOHN
```

Упражнение:

- Создайте функцию, возвращающую 2 целочисленных значения
- Функция принимает 2 целых числа (int) в качестве аргументов
- Первое возвращаемое значение — сумма аргументов, второе — их разность

```go
package main

import "fmt"

func operation() () {

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

func operation(x, y int) (int, int) {
    var sum, substraction int
    sum = x + y
    substraction = x - y
    return sum, substraction
}

func main () {
    // Ваш код здесь
    sum, subs := operation(10,5)
    fmt.Println(sum, subs)
}

```

</details>
