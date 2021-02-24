package main

import (
	"fmt"
	"time"
)

func Text_goroute(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	nums := 1000000
	fmt.Printf("nums:", nums)
	start := time.Now()
	for i := 0; i < nums; i++ {
		go Text_goroute(i, i*2)
	}
	end := time.Since(start)
	fmt.Println("TestGoroute;nums=", nums, "\n消耗时间=", end)
}
