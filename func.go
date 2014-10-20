package main


import "fmt"

func average(xs []int) int {    
    total := 0
    for _, v := range xs {
        total += v
    }
    return total / int(len(xs))
}

func main(){
 xs := []int{98,93,77,82,83}
fmt.Println(average(xs))
	fmt.Println("hi man function")
b:="sdafdasfdadsf"
fmt.Println(b)

	
}