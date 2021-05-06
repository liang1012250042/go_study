package main

import "fmt"

//定义函数类型
//为已存在的数据类型起别名
type Int int

//方法
//func (方法接受者)方法名(参数列表)返回值类型
func (a Int)add(b Int)Int{
	return a+b
}

func main() {
	var a Int =100
	//var b Int

	value:=a.add(20)
	fmt.Println(value)
}
