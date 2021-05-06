package main

import "fmt"

func main() {
	arr:=make([]interface{},3)
	arr=append(arr,"111",2,34,12,[]int{3,5,6},"强无敌")

	for _,v:=range arr{
		data,ok:=v.(int)
		if ok {
			fmt.Println(ok)
			fmt.Println(data)
		}
	}
}
