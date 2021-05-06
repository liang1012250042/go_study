package main

import "fmt"

type person struct {
	name string
	age int
	sex string
}

type student struct {
	//匿名字段
	person
	id int
	score int
}


func main() {
	var stu student
	stu.name = "匿名字段"
	fmt.Println(stu)
	fmt.Println("可以用..........")
}
