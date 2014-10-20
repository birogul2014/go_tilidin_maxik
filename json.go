package main
import "encoding/json"
import "fmt"
import "strings"
func main() {
	res:=map[string]int{"name":5,"uyghur":2}
	resb,_:=json.Marshal(res)
	fmt.Println(string(resb))
	  slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
	city := strings.SplitN("tormahiri/query?q=tokyo", "/", 3)[2]
	fmt.Println(city)
}