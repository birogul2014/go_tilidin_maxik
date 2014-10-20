package main

import"fmt"

func on(xptr *int){
		*xptr=1
	}
	
func main()
{

	 xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr) // x is 1
}