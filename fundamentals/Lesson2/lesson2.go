package main

import (
	"errors"
	"fmt"
)

func main() {

	greeting1()
	greeting2("Thành")
	fmt.Println(sum(5, 2))
	var1, var2, err := div(7, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else if var2 == 0 {
		fmt.Println(var1)
	} else {
		fmt.Println(var1, var2)
	}
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case var2 == 0:
		fmt.Println(var1)
	default:
		fmt.Println(var1, var2)
	}
	var3 := 20
	switch var3 {
	case 10:
		fmt.Println("A")
	case 20, 30:
		fmt.Println("B")
	case 40:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}

func greeting1() {
	fmt.Println("Hello, World!")
}

func greeting2(name string) {
	fmt.Println("Welcome, " + name)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("cannot div by 0")
		return 0, 0, err
	}
	return num1 / num2, num1 % num2, err
}
