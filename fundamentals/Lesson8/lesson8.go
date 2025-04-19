package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// func main() {

// 	// var c = make(chan int, 2)
// 	// c <- 1
// 	// fmt.Println(<-c)

// 	// var c = make(chan int)
// 	// go process(c)
// 	// fmt.Println(<-c)

// 	// var c = make(chan int)
// 	// go process_(c)
// 	// for i := 0; i <= 5; i++ {
// 	// 	fmt.Println(<-c)
// 	// }

// 	// var c = make(chan int)
// 	// go process__(c)
// 	// for i := range c {
// 	// 	fmt.Println(i)
// 	// }

// 	// var c = make(chan int)
// 	// go process___(c)
// 	// for i := range c {
// 	// 	fmt.Println(i)
// 	// }

// 	// var c = make(chan int)
// 	var c = make(chan int, 5)
// 	go process____(c)
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// 	fmt.Println("End")

// }

// func process(c chan int) {
// 	c <- 123
// }

// func process_(c chan int) {
// 	for i := 0; i <= 5; i++ {
// 		c <- i
// 	}
// }

// func process__(c chan int) {
// 	for i := 0; i <= 5; i++ {
// 		c <- i
// 	}
// 	close(c) // Nhiều go gọi tới khả năng là đưa dữ liệu vào channel bị đóng sẽ gặp lỗi
// }

// func process___(c chan int) {
// 	defer close(c) // Đảm bảo đống khi hàm kết thúc kế cả khi lỗi
// 	for i := 0; i <= 5; i++ {
// 		c <- i
// 	}
// }

// func process____(c chan int) {
// 	defer close(c)
// 	for i := 0; i <= 5; i++ {
// 		c <- i
// 	}
// 	fmt.Println("Done")
// }

var MAX_CHICKEN_PRICE = float64(5)
var MAX_TOFU_PRICE = float64(3)

func main() {

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"a", "b", "c"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float64() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float64() * 20
		if chickenPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select { // Nó sẽ kiểm tra xem là channel nào có biến động trước thì nó thực hiện case của channel đó
	case website := <-chickenChannel:
		fmt.Println("Chicken", website)
	case website := <-tofuChannel:
		fmt.Println("Tofu", website)
	}
}
