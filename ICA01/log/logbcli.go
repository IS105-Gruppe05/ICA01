package main


import "fmt"
import "os"
import "strconv"
//import "./log"
import "bufio"
import "./log"



func main() {
      scanner := bufio.NewScanner(os.Stdin)
      var text string
      var text2 string
      fmt.Println("Skriv q som tall og base for å slutte programmet")

      for text != "q"{
        fmt.Print("Skriv inn tall: ")
        scanner.Scan()
        text = scanner.Text()
        fmt.Print("Skriv base: ")
        scanner.Scan()
        text2 = scanner.Text()
        if text!= "q" {


          tall, err  := strconv.ParseFloat(text, 64)
                if err!=nil {
                        fmt.Println(err)
                        fmt.Println("Tallet kan ikke vere: " + text)
                }
          base, err := strconv.ParseFloat(text2, 64)
                if err!=nil {
                        fmt.Println(err)
                        fmt.Println("Basen kan ikke vere: " + text2)
                }


          fmt.Printf("That number is %f!\n", log.Logresult(tall, base))


        }
      }

}
