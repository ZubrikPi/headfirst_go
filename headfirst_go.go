// package main
// import "fmt"
// func main() {
//   fmt.Println("Hello, Go!")
// }
//
// package main
// import(
//   "fmt"
//   "reflect"
// )
// func main() {
//   fmt.Println(reflect.TypeOf(42))
//   fmt.Println(reflect.TypeOf(3.1415))
//   fmt.Println(reflect.TypeOf(true))
//   fmt.Println(reflect.TypeOf("Hello, Go!"))
// }
// int, float64, string - всё это ТИПЫ значений
// package main
// import "fmt"
// func main() {
//   var quantity int
//   var length, width float64
//   var customerName string
//
//   quantity = 4
//   length, width = 1.2, 2.4
//   customerName = "Damon Cole"
//
//   fmt.Println(customerName)
//   fmt.Println("has ordered", quantity, "sheets")
//   fmt.Println("each with an area of")
//   fmt.Println(length*width, "square meters")
// }

// package main
// import "fmt"
// func main() {
//   var quantity int
//   var length, width float64
//   var customerName string
//
//   quantity = 4
//   length, width = 1.2, 2.4
//   customerName = "Damon Cole"
//
//   fmt.Println(customerName)
//   fmt.Println("has ordered", quantity, "sheets")
//   fmt.Println("each with an area of")
//   fmt.Println(length*width, "square meters")
// }

// package main
// import "fmt"
// func main() {
//   var quantity int            = 4
//   var length, width float64   = 1.2, 2.4
//   var customerName string     = "Damon Cole"
//   //
//   // quantity = 4
//   // length, width = 1.2, 2.4
//   // customerName = "Damon Cole"
//   //
//   fmt.Println(customerName)
//   fmt.Println("has ordered", quantity, "sheets")
//   fmt.Println("each with an area of")
//   fmt.Println(length*width, "square meters")
// }

// package main
// import "fmt"
// func main() {
//   quantity          := 4
//   length, width     := 1.2, 2.4
//   customerName      := "Damon Cole"
//
//   fmt.Println(customerName)
//   fmt.Println("has ordered", quantity, "sheets")
//   fmt.Println("each with an area of")
//   fmt.Println(length*width, "square meters")
// }

// задание "Развлечения с магнитами"
// package main
// import "fmt"
// func main() {
//   var originalCount int = 10
//   var eatenCount    int = 4
//
//   fmt.Println("I starter with", originalCount, "apples")
//   fmt.Println("Some jerk ate", eatenCount, "apples")
//   fmt.Println("There are", originalCount-eatenCount, "apples left")
// }

// package main
// import "fmt"
// func main() {
//   var length float64 = 1.2
//   var width  int = 2       // из-за int программа не выполнится
//                            // так как значения должны относиться к
//                            // одному типу
//   fmt.Println("Area is", length*width)
//   fmt.Println("length > width?", length > width)
// }

// package main
// import "fmt"
// func main() {
//   var length float64 = 1.2
//   var width  float64 = 2       // теперь программа выполнится
//   fmt.Println("Area is", length*width)
//   fmt.Println("length > width?", length > width)
// }

// package main
// import "fmt"
// func main() {
//   var length float64 = 1.2
//   var width  int     = 2
//
//   fmt.Println("Area is", length*float64(width))
//   fmt.Println("length > width?", length > float64(width))
//}

package main
import "fmt"
func main() {
  var length float64 = 1.2
  var width  int = 2

  fmt.Println("Area is", length*float64(width))
  fmt.Println("length > width?", length > float64(width))
}

// test string
// darkblood

// package main
// import "fmt"
// func main() {
// 	fmt.Println("Hello world")
// }

// package main
// import "fmt"
// func main() {
//     var price float64 = 100
//
//     fmt.Println("Price is", price, "dollars")
//
//     var taxRate float64 = 0.08
//
//     var tax float64 = (float64(price) * taxRate)
//     fmt.Println("Tax is", tax, "dollars")
// }

package main
import "fmt"
func main() {
    var price int = 100

    fmt.Println("Price is", price, "dollars")

    var taxRate float64 = 0.08

    var tax float64 = (float64(price) * taxRate)
    fmt.Println("Tax is", tax, "dollars")
}

// ниже код из примера на habr
// package main
//
// import (
//   "fmt"
//   "strconv"
// )
//
// func main() {
//   var t map[string]interface{}     // Создаем переменную, сейчас она nil и мы не можем в неё класть ключи
//   t = make(map[string]interface{}) // Создаем map  в памяти, теперь можем класть ключи
//   t["key"] = "1"                   // Кладём ключ
//
//   var num int
//
//   rawValue, ok := t["key"] // Вытаскиваем ключ из map
//   if ok {
//     stringValue, ok := rawValue.(string) // Вытаскиваем string из interface{}
//     if ok {
//       var err error
//       num, err = strconv.Atoi(stringValue) // Парсим string в int
//       if err == nil {
//         fmt.Println(num)
//       } else {
//         fmt.Println("Число имеет неверный формат")
//       }
//     } else {
//       fmt.Println("Внутри interface{} лежит не тип string")
//     }
//   } else {
//     fmt.Println("Такого ключа нет в map")
//   }
// }
// 
