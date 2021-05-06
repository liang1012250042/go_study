package main

import (
	"errors"
	"fmt"
)

func test(a int,b int) (value int,err error) {
	if b==0 {
		err=errors.New("错误 除数不可为0")
		return
	}else {
		value=a/b
		return
	}
}

func main() {
	value,err:=test(10,0)
	//err不等于空 有错误
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(value)
	}
}
