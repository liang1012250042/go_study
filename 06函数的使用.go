package main

import "fmt"

func print(){
	fmt.Println("我是你爹")
}

func getNum(a int, b int) int {
	return a+b
}

type FUNCTYPE func()
type FUNCTEST func()

func main() {
	var _ FUNCTYPE
	P := print
	P()

}
