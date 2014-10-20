package main
import "fmt"

 func printCommand(cmd,cmds int) {
fmt.Println(cmd+cmds)
}
func main() {
 val:=make(chan string)

 
 go	func () {	val <-"kimmat"}()
 msg:=<-val
 fmt.Println(msg)
 kimmat:=make(chan string)

 go func(){kimmat<-"hay man channel kimmiti"}()
 uqur:=<-kimmat
 fmt.Println(uqur)
 type uqurlar struct{
 	wd,hg int
 }
 sq:=uqurlar{wd:30,hg:30}
 fmt.Println(sq)
 san:=make(chan int)
 go func(){san<-35}()
 sanlar:=<-san
 fmt.Println(sanlar)

printCommand(25,20)
}