package main



import (
"fmt"
"strings"
"bytes"
"container/list"
)

func main() {
	var x list.List
	x.PushBack(1)
	for e:=x.Front(); e!=nil;e=e.Next(){

		fmt.Println(e.Value)
	}
	arr := []byte("test")
	var buf bytes.Buffer
buf.Write(arr)
	fmt.Println(strings.Contains("test", "es"))
}