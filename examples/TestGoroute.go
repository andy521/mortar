package main

import (
	"fmt"
	"time"
)

func Text_goroute(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func ff(v ...interface{}) {
	fmt.Println(v)
}

/**
普通的调用方法
*/
func main() {
	nums := 100000
	fmt.Printf("nums:", nums)
	start := time.Now()
	for i := 0; i < nums; i++ {
		go ff([]interface{}{i, i * 2, "hello"})
		//go Text_goroute(i, i*2)
	}
	end := time.Since(start)
	fmt.Println("TestGoroute;nums=", nums, "\n消耗时间=", end)
}
