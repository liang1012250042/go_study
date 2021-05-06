package main

import "fmt"

type person2 struct {
	name string
	age int
}

type student2 struct {
	person2
	score int
}

func (p *person2)sayHello()  {
	fmt.Println("父类方法")
}

func (s student2)sayHello()  {
	fmt.Println("子类方法")
}

//接口定义
type humaner interface {
	//方法的声明 函数声明 没有具体实现 具体实现要根据对象的不同 实现方式也不同
	sayHello()
}

func main() {
	var stu student2=student2{person2{"小hong",66},456}
	var h humaner
	h=&stu
	h.sayHello()
}
