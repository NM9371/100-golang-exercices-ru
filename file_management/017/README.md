# Упражнение: Чтение файла

Для чтения файла используется функция [os.ReadFile()](https://pkg.go.dev/os#ReadFile).
Её сигнатура:

```go
func ReadFile(name string) ([]byte, error)
```

Обратите внимание, что она возвращает массив `byte`. Как его читать людям? — спросите вы.
Вот здесь мы познакомимся с преобразованием типов.

## Преобразование типов

Преобразование типов — это процесс преобразования одного типа в другой. Например, `int32` в `float64` или `[]byte` в `string` и обратно.

Не путать с приведением типов (type casting) — в Go его нет.
> При приведении типов программист явно преобразует тип с помощью оператора приведения. При преобразовании типов это делает компилятор.

## Как преобразовать тип

В Go для преобразования типа нужно указать целевой тип и обернуть переменную в скобки:

```go
var i int = 42
var f float64 = float64(i) // теперь f содержит значение 42, но уже не целое число!
```

Помните, что Go — язык со статической типизацией! Если попытаться напрямую присвоить переменную `i` (типа `int`) в `f` (типа `float64`), получите ошибку `cannot use i (variable of type int) as float64 value in variable declaration`!

Упражнение: Прочитайте файл `read.txt` в этой директории.

Подсказка: Массив байт можно преобразовать в строку с помощью `string(myArrayOfBytesVariable)`

```go
package main

import (
  "log"
  "fmt"
)

func main() {
  data, err := os.ReadFile()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println()
}
```

<details>
<summary> Решение: </summary>

```go
package main

import (
  "log"
  "fmt"
)

func main() {
  data, err := os.ReadFile("/tmp/test.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(data)) // или os.Stdout.Write(data)
}
```

</details>
