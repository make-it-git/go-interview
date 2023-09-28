package main

import (
	"fmt"
	"runtime"
)

func slices() {
	s := getSubSlice()
	printMemStat()
	runtime.GC()
	for i := 1; i < 10; i++ {
		s2 := getSubSlice()
		runtime.GC()
		printMemStat()
		fmt.Println(s2)
	}
	fmt.Println(s)
	runtime.GC()
	printMemStat()
}

func getSubSlice() []int {
	s := make([]int, 1_000_000)

	return s[999_998:]
}

func printMemStat() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 1024 / 1024)
}
