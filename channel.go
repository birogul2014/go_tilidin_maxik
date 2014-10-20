package main

import "fmt"

func main(){
   message:=make(chan string)

   go func () {message <- "hello tormahirid"}()
   msg:=<-message
   fmt.Println(msg)
}