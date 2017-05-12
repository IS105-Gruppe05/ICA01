package main

import "fmt"
import "os"
import "strconv"
import "./log"


func main() {
number, err := strconv.ParseFloat(os.Args[1], 64)
if err != nil {fmt.Println("Arguments are Log(number) Log(base)", err)
} else {
    fmt.Printf("That number is %f!", log.Logresult(number))
  }
}
