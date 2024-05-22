### Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.
```go
package main

import (
  "fmt"
)

func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}

func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}

Что выведет программа? Объяснить вывод программы.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func asChan(vs ...int) <-chan int {
   c := make(chan int)

   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }

      close(c)
  }()
  return c
}

func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}

func main() {

   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}

```
### Ответ 
```
Программа выведет числа от 1 до 8 включительно в произвольном порядке и после этого будет
бесконечно выводить число 0.
функция asChan принимает значения для записи в созданный канал и возвращает его, 
после того как значения закончатся, канал закрывается
функция merge принимает 2 канала для записи в созданный канал и возвращает его,
то есть служит для объединения 2 каналов в 1
После этого производится чтение происходит вывод из объединенного канала,
но тк нет проверки на закрытие канала в функции merge, она продолжает туда 
писать дефолтные значения типа данных канала, в нашем случае это 0
```
