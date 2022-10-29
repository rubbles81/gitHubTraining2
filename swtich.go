package main
import ("fmt")

func main() {
   day := "cat"

   switch day {
   case "cat","dog","mose":
    fmt.Println("Odd weekday")
   case "re","we":
     fmt.Println("Even weekday")
   case "ee","eeee":
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}
