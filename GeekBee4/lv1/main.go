package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanln(&x)
		mp[x]++
	}
	for key, value := range mp {
		println("key:", key, " value:", value, "\n")
	}
}
