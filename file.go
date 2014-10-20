package main


import (
    "os"
   "bufio"
   "fmt"
)

}
func main() {
	  f , err := os.Open("C:/a.txt",os.O_RDONLY,0)
   if err != nil{
       fmt.Printf("%v\n",err)
   }
       
   
}