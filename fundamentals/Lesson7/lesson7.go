package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.Mutex{}
var rwm = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var result = []string{}

func main() {

	// t0 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	dbCall(i)
	// }
	// fmt.Println(time.Since(t0))

	// t0 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1)
	// 	go dbCall(i)
	// }
	// wg.Wait()
	// fmt.Println(time.Since(t0))

	// t0 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1)
	// 	go dbCall_(i)
	// }
	// wg.Wait()
	// fmt.Println(time.Since(t0))
	// fmt.Println(result)

	// t0 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1)
	// 	go dbCall__(i)
	// }
	// wg.Wait()
	// fmt.Println(time.Since(t0))
	// fmt.Println(result)

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall___(i)
	}
	wg.Wait()
	fmt.Println(time.Since(t0))
	fmt.Println(result)

}

func dbCall(i int) {
	var delay float64 = rand.Float64() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(dbData[i])
	wg.Done()
}

func dbCall_(i int) {
	var delay float64 = 1
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(dbData[i])
	result = append(result, dbData[i]) // Dễ bị missing data
	wg.Done()
}

func dbCall__(i int) {
	var delay float64 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(dbData[i])
	m.Lock()
	result = append(result, dbData[i])
	m.Unlock()
	wg.Done()
}

func dbCall___(i int) {
	var delay float64 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(value string) {
	rwm.Lock()
	result = append(result, value)
	rwm.Unlock()
}

func log() {
	rwm.RLock()
	fmt.Println(result)
	rwm.RUnlock()
}
