package main

import "fmt"

func main() {
	var a [] int
	fmt.Println("a:",a)

	//创建切片
	b:=make([]int,5)
	//通过下标对切片进行赋值
	b[0]=123
	b[1]=123
	b[2]=123
	b[3]=123
	b[4]=123
	//动态申请 继续添加 不会越界
	b=append(b,666)
	b=append(b,666,777,888)
	fmt.Println("b:",b)

	//切片截取
	c:=[]int{1,2,3,4,5}
	slice1:=c[0:]
	slice2:=c[3:]
	slice3:=c[3:4]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	//切片的追加和拷贝
	d:=make([]int,5)
	d=append(d,666)
	fmt.Println("d:",d)
	e:=make([]int,5)
	copy(e,d)
	fmt.Println("e:",e)
}
