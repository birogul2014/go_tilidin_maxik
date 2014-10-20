package main

import "fmt"
func f(){
	fmt.Println("helo")
}
func s(d int){
	fmt.Println(d)
}
type person struct{
	name int
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
type android struct{
	person
}
func main() {
     a:=new(android)
     a.person.Talk()
	var c person
	c.name=20

	var x int =5
	s(c.name+x)

	f()
}