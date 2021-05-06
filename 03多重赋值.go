package main

import "fmt"

func main() {
	a,b,c,d:=10,20,30,40.0
	e:=100
	f:=200
	fmt.Println(e,f)
	e,f=f,e
	fmt.Println(e,f)
	fmt.Println(a,b,c,d)
}
