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
//   var lenght, width float64
//   var customerName string
//
//   quantity = 4
//   lenght, width = 1.2, 2.4
//   customerName = "Damon Cole"
//
//   fmt.Println(customerName)
//   fmt.Println("has ordered", quantity, "sheets")
//   fmt.Println("each with an area of")
//   fmt.Println(lenght*width, "square meters")
// }

package main
import "fmt"
func main() {
  var quantity int
  var lenght, width float64
  var customerName string

  quantity = 4
  lenght, width = 1.2, 2.4
  customerName = "Damon Cole"

  fmt.Println(customerName)
  fmt.Println("has ordered", quantity, "sheets")
  fmt.Println("each with an area of")
  fmt.Println(lenght*width, "square meters")
}
