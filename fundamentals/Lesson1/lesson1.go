package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	// fmt.Println("Hello, World!")

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	// var num1 int
	// fmt.Println(num1)
	// var num2 int = 100
	// fmt.Println(num2)
	// var num3 float32 = 12345678.9
	// fmt.Printf("%f\n", num3)
	// var num4 float64 = 12345678.9
	// fmt.Printf("%f\n", num4)

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	var num1 float32 = 10.1
	var num2 int32 = 2
	var num3 float32 = num1 + float32(num2)
	var num4 int32 = int32(num1) + num2
	fmt.Println(num3)
	fmt.Println(num4)

	var num5 int = 7
	var num6 int = 2
	fmt.Println(num5 / num6)
	fmt.Println(num5 % num6)
	fmt.Println(float32(num5) / float32(num6))

	var str1 string = "Hello, World!\nWelcome to Go!"
	var str2 string = `Goodbye, World!
See you again, Go!`
	fmt.Println(str1)
	fmt.Println(str2)

	var str3 string = "Hello, World!"
	fmt.Println(len(str3))
	fmt.Println(utf8.RuneCountInString(str3))

	var rune1 rune = 'A'
	fmt.Println(rune1)

	var bool1 bool = false
	var bool2 bool = 1 < 2
	fmt.Println(bool1)
	fmt.Println(bool2)

	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)
	fmt.Println(var1, ", ", var2)
	fmt.Println(strconv.Itoa(var1) + ", " + strconv.Itoa(var2))

	var3 := 100
	var4, var5 := 1, 2
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)

	var6 := 1
	var7 := "abc"
	var8 := 'A'
	fmt.Printf("%v %v %v\n", var6, var7, var8)

	// const const1 float64 // Lỗi!
	const const1 float64 = 3.14
	// const1 = 10 // Lỗi!
	fmt.Println(const1)

}
