# Упражнение: Две строковые переменные

Со строками можно выполнять операции. Например, можно составить фразу, сложив 2 строковые переменные:

```go
fmt.Println("My pet is called " + dog_name + ", and he loves" +  dog_favourite_food)
```

Существует несколько [соглашений](https://en.wikipedia.org/wiki/Naming_convention_(programming)#Go) об именовании переменных: `thisVariable` (camelCase), `this_variable` (snake_case), `this-variable` (kebab-case) и другие.

В golang предпочтительны camelCase и UpperCamelCase; написание первой буквы в верхнем регистре экспортирует этот фрагмент кода, а нижний регистр делает его доступным только в текущей области видимости.

Упражнение: Конкатенируйте две строковые переменные и выведите результат

- Одна переменная должна называться "helloWorld", другая — "itsMeMario"
- Между ними должен быть пробел

```go
package main

func main() {
  // Здесь пишите ваш код
  var helloWorld, itsMeMario string
  // ...
}
```

<details>
<summary> Решение: </summary>

```go
package main

import "fmt"

func main() {
  // Создаём новую переменную с именем helloWorld
  var helloWorld, itsMeMario string
  helloWorld = "Hello World!"
  itsMeMario = "It's a me, Mario"
  // Выводим переменную
  fmt.Println(helloWorld + " " + itsMeMario)
}

// Чтобы запустить программу:
// - go run solution.go
```

</details>
