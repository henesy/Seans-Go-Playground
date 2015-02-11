package main
import (
  "fmt"
  "time"
  "math/rand"
  "C"
)
func main(){
  rand.Seed( time.Now().UTC().UnixNano())
  a := 1 + rand.Intn(7-1)
  rand.Seed( time.Now().UTC().UnixNano())
  b := 1 + rand.Intn(7-1)
  if a == b {
    fmt.Print("You die\n")
  } else {
    fmt.Print("You live\n")
  }
}
