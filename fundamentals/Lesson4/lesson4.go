package main

import (
	"fmt"
	"strings"
)

func main() {

	// // var str1 string = "Thành"
	// str1 := "Thành"
	// fmt.Println(str1)
	// var indexed = str1[0]
	// fmt.Printf("%v, %T\n", indexed, indexed)
	// for i, v := range str1 {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(len(str1))

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	// var rune1 []rune = []rune("Thành")
	// fmt.Println(rune1)
	// var indexed = rune1[0]
	// fmt.Printf("%v, %T\n", indexed, indexed)
	// for i, v := range rune1 {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(len(rune1))

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	// var slice []string = []string{"T", "h", "à", "n", "h"}
	// var str string = ""
	// for i := range slice {
	// 	str += slice[i]
	// }
	// fmt.Println(str)
	// // str[0] = "asd" // Lỗi

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	var slice []string = []string{"T", "h", "à", "n", "h"}
	var strBuilder strings.Builder
	for i := range slice {
		strBuilder.WriteString(slice[i])
	}
	fmt.Println(strBuilder)
	var str string = strBuilder.String()
	fmt.Println(str)

}
