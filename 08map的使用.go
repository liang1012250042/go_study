package main

import "fmt"

func mapTest(m map[string]string)  {
	m["333"] = "789"
}

func main()  {

	m:=make(map[int]string,1)
	m[122] = "我又是你d"

	n:=make(map[string]string,1)
	n["123"] = "test"
	n["1234"] = "tes1t"

	for k,v:=range n{
		fmt.Println(k,v)
	}

	o:=map[string]string{"1":"123","2":"456"}

	fmt.Println(o)

	delete(o,"1")

	fmt.Println(o)

	mapTest(o)

	fmt.Println(o)

	fmt.Println(n)
	fmt.Println("我是你d")
}
