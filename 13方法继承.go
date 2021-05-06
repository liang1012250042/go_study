package main

import "fmt"

type person1 struct {
	id int
	name string
	age int
}

type student1 struct {
	person1
	score int
}

func (p person1)sayHello(){
	fmt.Println("so what")
}

func main() {
	var stu1 student1
	stu1.name="聂cx"
	stu1.age=22
	stu1.id=1
	stu1.score=11
	stu1.sayHello()
}
