package main


import (
"fmt"

"os"
"time"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main(){
	var f:=8.2
	var fn:=Sqrt(f)
       
	  fmt.Println(time.Now())

}