# Упражнение: Переименование файла

Это простое упражнение — вы заслужили!
Пакет "os" также содержит функцию для [переименования](https://pkg.go.dev/os#Rename) файла (при наличии прав доступа и существования файла).

Сигнатура функции `os.Rename`:

```go
func Rename(oldpath, newpath string) error
```

Упражнение: Переименуйте существующий файл — смените старое имя на новое.

```go
package main

import "os"


func main () {
  var old_name, new_name string
  // Ваш код здесь

}
```

<details>
<summary> Решение: </summary>

```golang
package main

import "os"


func main () {
  var src, dest string
  // Ваш код здесь

  src = "name1.txt"
  dest = "name2.txt"

  os.Rename(src, dest)
}
```

</details>
