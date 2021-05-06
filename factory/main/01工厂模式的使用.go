package main

import (
	"fmt"
	"test/factory/model"
)

func main() {
	//var stu = model.Student{
	//	Name: "liang",
	//	Score: 78.9,
	//}

	//fmt.Println(stu)

	//单要导入的类是小写(私有变量),我们可以使用工厂模式来解决
	//var p = model.NewPerson("liang1027", 1999.1027)
	//
	//fmt.Println(*p)

	var p = model.NewPerson("liang1027", 1999.1027)

	fmt.Println(*p)
	fmt.Println("name",p.Name,"score",p.GetScore())



}
