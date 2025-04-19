package main

import "fmt"

type student struct {
	id   string
	name string
}

func main() {

	// var p *int32
	// var i int32
	// fmt.Println(p)
	// fmt.Println(i)

	// var p *int32 = new(int32) // Cấp phát một vùng nhớ cho kiểu int32 với giá trị mặc định là 0 và địa chỉ vùng nhó đó sẽ được gán cho p (tức p có giá trị là địa chỉ của ô nhớ lưa trữ số 0 vừa được cấp phát)
	// // p := new(int32)
	// fmt.Println(p)
	// fmt.Println(*p)
	// *p = 10 // Thay đổi giá trị của ô nhớ mà p đang trỏ tới
	// fmt.Println(p)
	// fmt.Println(*p)

	// var p *int32 = new(int32)
	// var i int32
	// fmt.Println(p, *p)
	// fmt.Println(i)
	// p = &i // Trỏ đến địa chỉ của vùng nhớ mà i mang giá trị
	// *p = 100
	// fmt.Println(p, *p)
	// fmt.Println(i)
	// // var k int32 = i // Đây chỉ là copy giá trị của i vào một ô nhớ mới cho k thôi chứ không hoạt động như con trỏ

	// var slice []int32 = []int32{1, 2, 3}
	// var sliceCopy = slice // Tham chiếu cùng
	// sliceCopy[1] = 100
	// fmt.Println(slice)
	// fmt.Printf("%p\n", &slice)
	// for i := range slice {
	// 	fmt.Printf("%p ", &slice[i])
	// }
	// fmt.Println()
	// fmt.Println(sliceCopy)
	// fmt.Printf("%p\n", &sliceCopy)
	// for i := range sliceCopy {
	// 	fmt.Printf("%p ", &sliceCopy[i])
	// }
	// fmt.Println()

	// var arr [3]int32 = [3]int32{1, 2, 3}
	// var arrCopy = arr // Copy
	// arrCopy[1] = 100
	// fmt.Println(arr)
	// fmt.Printf("%p\n", &arr)
	// for i := range arr {
	// 	fmt.Printf("%p ", &arr[i])
	// }
	// fmt.Println()
	// fmt.Println(arrCopy)
	// fmt.Printf("%p\n", &arrCopy)
	// for i := range arrCopy {
	// 	fmt.Printf("%p ", &arrCopy[i])
	// }
	// fmt.Println()

	// var slice []int32 = []int32{1, 2, 3}
	// var sliceCopy = make([]int32, len(slice))
	// copy(sliceCopy, slice)
	// sliceCopy[1] = 100
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	// var arr1 [5]float64 = [5]float64{1, 2, 3, 4, 5}
	// var arr2 [5]float64 = square(arr1) // Copy
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// var slice1 []float64 = []float64{1, 2, 3, 4, 5}
	// var slice2 []float64 = square_(slice1)
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// var arr1 [5]float64 = [5]float64{1, 2, 3, 4, 5}
	// var arr2 [5]float64 = square__(&arr1) // Tham chiếu cùng
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	var stu1 = student{"01", "Thành"}
	// var stu2 = stu1 // Copy
	var stu2 = new(student)
	stu2 = &stu1
	stu2.id = "XXX"
	fmt.Println(stu1)
	fmt.Println(*stu2)

}

func square(arr [5]float64) [5]float64 {
	for i := range arr {
		arr[i] *= arr[i]
	}
	return arr
}

func square_(slice []float64) []float64 {
	for i := range slice {
		slice[i] *= slice[i]
	}
	return slice
}

func square__(arr *[5]float64) [5]float64 {
	for i := range arr {
		arr[i] *= arr[i]
	}
	return *arr
}
