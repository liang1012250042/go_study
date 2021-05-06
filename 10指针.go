package main

import "fmt"

func pointerTest(a *int,b *int)  {
	*a,*b = *b,*a
}

func main() {

	//a:=10
	//
	//var p *int
	//p=&a
	//
	//fmt.Println(*p)
	//fmt.Println("又tmd是指针")
	//
	//var p *int
	//var q *string
	//p=new(int)
	//q=new(string)
	//fmt.Println(p)
	//fmt.Println(*p)
	//fmt.Println("---------")
	//fmt.Println(q)
	//fmt.Println(*q)
	//
	//
	//a:=10
	//b:=20
	//fmt.Println(a)
	//fmt.Println(b)
	//
	//pointerTest(&a,&b)
	//fmt.Println(a)
	//fmt.Println(b)

	//var arr [5]int =[5]int{1,2,3,4,5}
	//var p *[5]int
	//p= &arr

	var arr []int = []int{1,2,3,4,5}
	p:=&arr
	(*p)[1]  = 3
	fmt.Println(arr)

}
