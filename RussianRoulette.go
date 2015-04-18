package main
import (
  "fmt"
  "time"
  "math/rand"
)
func main(){
  rand.Seed( time.Now().UTC().UnixNano())
  if b := 1 + rand.Intn(7-1); 1 == b {
    fmt.Print("You die\n")
  } else {
    fmt.Print("You live\n")
  }
}
