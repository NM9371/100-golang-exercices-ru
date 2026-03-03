# Упражнение: Массивы II

Строковая переменная — это массив более мелких элементов типа `char`.

Например, слово `TREE` — это массив из 4 символов, каждому из которых присвоен индекс.
Первый элемент (с индексом 0) — буква 'T', второй (с индексом 1) — буква 'R' и так далее.

```txt
T - R - E - E
0 - 1 - 2 - 3
```

Упражнение: Используя единственную строковую переменную с именем helloWorld, выведите только первый символ.

- Для доступа к значению используйте `string[n]`, где n — номер индекса.

```go
package main

import "fmt"

func main () {
  var helloWorld string
  helloWorld = "Hello World!"
  // Здесь пишите ваш код
  fmt.Println()
}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

func main() {
  // Создаём новую переменную с именем helloWorld
  var helloWorld string
  helloWorld = "Hello World!"
  // Выводим первую букву
  fmt.Println(helloWorld[0])
}

// Чтобы запустить программу:
// - go run solution.go
```

</details>
