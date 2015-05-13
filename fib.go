package main

import "fmt"

/* inefficient capped fibonacci */

func main() {
   fibs := make([]float32, 100)
   fibs[0]=0.0
   fibs[1]=1.0
   fmt.Printf("%3d: %.0f\n", 1, fibs[0])
   fmt.Printf("%3d: %.0f\n", 2, fibs[1])
   for i:=2;i<len(fibs);i+=1 {
       nums := fibs[i-2:i]
       fibs[i] = nums[0] + nums[1]
       fmt.Printf("%3d: %.0f\n", i+1, fibs[i])
   }
}
