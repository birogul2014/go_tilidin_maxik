package main
import "fmt"
import "sort"
import "os"
func main() {
strs :=[]string{"a","c","h"}
sort.Strings(strs)
fmt.Println(strs)
san:=[]int{1,2,5,76,3,4,15,343,12,44}
sort.Ints(san)
fmt.Println("san turlanid",san)
os.Setenv("A","taz")
fmt.Println("get kimmat",os.Getenv("A"))
s:=make([]string,3)
s[0]="salam"
s[1]="sadirdin"
s[2]="tormahiri"
fmt.Println("sliced data",s[0],s[1])
}
