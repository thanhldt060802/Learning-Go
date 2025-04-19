package main

import (
	"fmt"
	"time"
)

func main() {

	// var arr1 [3]int
	// arr1[1] = 123
	// fmt.Println(arr1[0])
	// fmt.Println(arr1[1:])
	// fmt.Println(arr1[1:3])
	// fmt.Println(arr1)
	// fmt.Println(&arr1[0])
	// fmt.Println(&arr1[1])
	// fmt.Println(&arr1[2])

	// var arr2 [3]int = [3]int{1, 2, 3}
	// var arr3 [3]int = [3]int{1, 2, 3}
	// arr4 := [...]int{1, 2, 3}
	// fmt.Println(arr2)
	// fmt.Println(arr3)
	// fmt.Println(arr4)

	// var slice1 []int = []int{1, 2, 3}
	// fmt.Println(slice1, len(slice1), cap(slice1))
	// slice1 = append(slice1, 10)
	// fmt.Println(slice1, len(slice1), cap(slice1))
	// var slice2 []int = []int{-1, -2}
	// slice1 = append(slice1, slice2...)
	// fmt.Println(slice1)

	// var slice3 []int = make([]int, 5, 8)
	// fmt.Println(slice3, len(slice3), cap(slice3))

	// var map1 map[string]int = make(map[string]int)
	// var map2 map[string]int = map[string]int{
	// 	"Thanh": 23,
	// 	"Phuc":  20,
	// }
	// fmt.Println(map1)
	// fmt.Println(map2)
	// fmt.Println(map2["Thanh"])
	// var1, var2 := map2["Phuc"]
	// fmt.Println(var1, var2)
	// var1, var2 = map2["Tam"]
	// fmt.Println(var1, var2)
	// delete(map2, "Thanh")
	// fmt.Println(map2)

	// var map3 map[string]int = map[string]int{
	// 	"Thanh": 22,
	// 	"Phuc":  19,
	// 	"Tam":   46,
	// 	"Hoa":   46,
	// }
	// for var3, var4 := range map3 {
	// 	fmt.Println(var3, var4)
	// }

	// var arr5 = [...]int{2, 4, 6, 8}
	// for var5, var6 := range arr5 {
	// 	fmt.Println(var5, var6)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i, "")
	// }
	// fmt.Println()
	// var i1 int = 0
	// for i1 < 10 {
	// 	fmt.Print(i1, "")
	// 	i1++
	// }
	// fmt.Println()
	// i1 = 0
	// for {
	// 	if i1 >= 10 {
	// 		break
	// 	}
	// 	fmt.Print(i1, "")
	// 	i1++
	// }
	// fmt.Println()

	// --------------------------------
	fmt.Println("--------------------------------")
	// --------------------------------

	var n int = 1000000
	var slice1 []int = []int{}
	var slice2 []int = make([]int, 0, n)

	fmt.Println(timeLoop(slice1, n))
	fmt.Println(timeLoop(slice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
