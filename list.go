package main


import (
"container/list"
"fmt"
)


func main(){

ice2 := slice[5:10]
  l := list.New()
  l.PushBack(1)
  l.PushBack(2)
  l.PushBack(3)
  l.PushBack(4)


  //遍历list，删除元素
  for e := l.Front(); e != nil; e = e.Next() {
      fmt.Println("removing", e.Value)
      l.Remove(e)
  }
  fmt.Println("After Removing...")
}