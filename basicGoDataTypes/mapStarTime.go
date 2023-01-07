package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var N = 40000000
	myMap := make(map[int]*int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	start := time.Now()
	runtime.GC()
	_ = myMap[0]

	duration := time.Since(start)
	fmt.Println("It took GC", duration, "to finish.")
}
