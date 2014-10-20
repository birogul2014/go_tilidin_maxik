package main
import(
"fmt"
"time"
)
type rec struct{
		width,height int
	}
	func(r *rec)rayun() int{
         return r.width*r.height
	}
func main() {
	ticker:=time.NewTicker(time.Millisecond *500)
	go func() {
	for t:=range ticker.C{
		fmt.Println("ticker at",t)
	}	
	}()
	time.Sleep(time.Millisecond *1500)
	

	
	result:=rec{width:10,height:5}
	fmt.Println(result.rayun())
	channel:=make(chan string,3)
	channel<-"adfadf"
	channel<-"dsfadsafsaf"
	for elm:=range channel{
		fmt.Println(elm)
	}
	close(channel)
	fmt.Println(channel)
	
}