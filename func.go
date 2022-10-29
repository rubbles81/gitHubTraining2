package main
import ("fmt")

func myFunction(x int, y string) (txt1 string, result txt1) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}
