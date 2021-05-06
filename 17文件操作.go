package main

import (
	"fmt"
	"os"
)

func main() {
	//file,err:=os.Create("./test.txt")
	//if err !=nil{
	//	fmt.Println("文件创建失败")
	//	return
	//}
	//
	//fmt.Println("文件创建成功")
	//
	//
	//
	//file.WriteString("塞林木\n")
	//file.WriteString("几百")
	//file.Close()

	file1,err:=os.OpenFile("./test.txt",os.O_RDWR,7)
	if err !=nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file1.Close()
	b:=make([]byte,1024)
	file1.Read(b)
		//io.WriteString(file1,"牛逼")
		//n,_:=file1.Seek(0,2)
		//file1.WriteAt([]byte("\n111"),n)
	fmt.Println(string(b))



}