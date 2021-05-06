package main
//
//import "fmt"
//
//type student struct {
//	stid int
//	name string
//	age float64
//}
//
//func stuctTest(s student)  {
//	fmt.Println(s)
//}
//
//func main()  {
//	fmt.Println("又见结构体...")
//	var stu1 student
//	stu1.stid = 1
//	stu1.name = "结构体测试1"
//	stu1.age = 666
//
//	stu2:=student{
//		stid: 66,
//		name: "结构体",
//		age: 3.14,
//	}
//
//	fmt.Println(stu1)
//	fmt.Println(stu2)
//
//	stu3:=make([]student,3)
//	stu3[0] = student{stid: 1,name: "1",age: 1}
//	stu3[1] = student{stid: 2,name: "2",age: 2}
//	stu3[2] = student{stid: 3,name: "3",age: 3}
//	fmt.Println(stu3)
//	stu3 = append(stu3,student{stid: 4,name: "4",age: 4})
//	fmt.Println(stu3)
//
//	stu4:=make(map[string]student)
//	stu4["1"] = student{stid: 1,name: "1",age: 1}
//	stu4["2"] = student{stid: 2,name: "2",age: 2}
//	fmt.Println(stu4)
//	delete(stu4,"1")
//	fmt.Println(stu4)
//
//	stuctTest(stu1)
//}
