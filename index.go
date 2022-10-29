package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idxy, valin := range fruits {
     fmt.Printf("%v\t%v\n", idxy, valin)
  }
}
