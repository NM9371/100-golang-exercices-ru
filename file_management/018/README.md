# Упражнение: Запись в файл

В этом упражнении мы создадим файл и запишем в него данные с помощью функции [os.Create()](https://pkg.go.dev/os#Create).

```go
// Сигнатура функции os.Create
func Create(name string) (*File, error)
```

Она возвращает указатель на тип `File`. Затем можно использовать функцию [os.WriteString()](https://pkg.go.dev/os#File.WriteString), чтобы записать строку в созданный файл.

```go
// Сигнатура функции os.WriteString
func (f *File) WriteString(s string) (n int, err error)
```

## Методы

Метод — это функция с особым параметром-получателем (receiver).

В сигнатуре выше мы видим часть `(f *File)` между ключевым словом `func` и именем функции.

Это означает, что функция `WriteString` может вызываться только на переменных типа `*File`. Получатель функции — `f` в сигнатуре выше.

Другой пример:

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
  // получатель функции Abs() — Vertex.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

В этом случае `Abs()` вызывается на переменной типа `Vertex` — `v`.

Упражнение: Запишите список из 5 стран в файл
Подсказка: `"\n"` можно использовать в строке для перевода строки.

```go
package main

import "os"


func main () {
  // Ваш код здесь

}
```

<details>
<summary> Решение: </summary>

```go
package main

import "os"

func main () {
  // Ваш код здесь
  file, err := os.Create("countries.txt")

  if err != nil {
    return
  }

  defer file.Close()

  file.WriteString("Germany\nFrance\nUSA\nSpain\nUK\n")
}
```

</details>
