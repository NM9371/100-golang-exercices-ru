# Упражнение: Range

Range — это идиоматический способ итерации по структурам данных; он широко используется для перебора массивов, отображений (map) или срезов (slice).

Синтаксис range:

```go
for index, value := range <структура_данных> {
  // index и value заполняются автоматически для каждого элемента структуры данных.
}
```

- Используйте range для вывода значений и индексов массива

```go
package main

import "fmt"

func main () {
  // инициализированный массив из 10 значений int [1..10]
  var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
  // Здесь пишите ваш код

}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

func main () {

  var arr = [10]int{1,2,3,4,5,6,7,8,9,10}

  for index, value := range arr {
    fmt.Print(index , ") " , value, "\n")
  }
}
```

</details>
