package main

import "fmt"

type student struct {
	id string
	name
	age int
	maj major
	string
}

type name struct {
	fName string
	lName string
}

type major struct {
	name string
	desc string
}

func (stu student) howAge() string {
	if stu.age <= 18 {
		return "young"
	} else {
		return "old"
	}
}

type master struct {
	id   string
	name string
	age  int
}

func (mas master) howAge() string {
	if mas.age <= 18 {
		return "young"
	} else {
		return "old"
	}
}

type person interface {
	howAge() string
}

func doSomething(per person) {
	fmt.Println(per.howAge())
}

func main() {

	// var stu1 student
	// fmt.Println(stu1)
	// fmt.Println(stu1.id)
	// fmt.Println(stu1.name)
	// fmt.Println(stu1.name.fName) // Không nhất thiết
	// fmt.Println(stu1.fName)
	// fmt.Println(stu1.age)
	// fmt.Println(stu1.maj)
	// fmt.Println(stu1.maj.name)
	// fmt.Println(stu1.string)

	// var stu2 student = student{id: "01", name: name{fName: "Lê", lName: "Thành"}, age: 23, maj: major{name: "IT", desc: "CS"}, string: "kaka"}
	// fmt.Println(stu2)
	// fmt.Println(stu2.id)
	// fmt.Println(stu2.name)
	// fmt.Println(stu2.fName)
	// fmt.Println(stu2.age)
	// fmt.Println(stu2.maj)
	// fmt.Println(stu2.maj.name)
	// fmt.Println(stu2.string)
	// stu2.name.lName = "Phúc"
	// fmt.Println(stu2)

	// var stu3 student = student{"02", name{"Lê", "Tâm"}, 46, major{"IT", "CS"}, "kaka"}
	// fmt.Println(stu3)

	// var per = struct {
	// 	id   string
	// 	name string
	// }{"01", "Thành"}
	// fmt.Println(per)

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	var stu student = student{"02", name{"Lê", "Tâm"}, 46, major{"IT", "CS"}, "kaka"}
	fmt.Println(stu.id)
	fmt.Println(stu)
	fmt.Println(stu.howAge())
	doSomething(stu)

	var mas person = master{id: "01", name: "Thành", age: 23}
	// fmt.Println(mas.id) // Lỗi!
	fmt.Println(mas.howAge())
	doSomething(mas)

}
