package main

import "fmt"

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  string
	Amount string
}

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpwkh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {

	// var slice1 = []int{1, 2, 3, 4, 5}
	// var slice2 = []float32{1.1, 2.2, 3.3}
	// fmt.Println(sumSlice(slice1))
	// fmt.Println(sumSlice(slice2))

	// var slice1 = []int{1, 2, 3, 4, 5}
	// var slice2 = []float32{}
	// fmt.Println(isEmpty(slice1))
	// fmt.Println(isEmpty(slice2))

	// var contacts []contactInfo = loadJSON[contactInfo]("/.contactInfo.json")
	// fmt.Printf("\n%+v", contacts)
	// var purchases []purchaseInfo = loadJSON[purchaseInfo]("/.purchaseInfo.json")
	// fmt.Printf("\n%+v", purchases)

	var myCar = car[gasEngine]{
		carMake:  "A",
		carModel: "B",
		engine: gasEngine{
			gallons: 3.14,
			mpg:     100,
		},
	}
	fmt.Println(myCar)

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
// 	data, _ = os.ReadFile(filePath)

// 	var loaded = []T{}
// 	json.Unmarshal(data, &loaded)

// 	return loaded
// }
