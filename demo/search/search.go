package main

import "sort"

func main() {

	s := []int{10, 20, 30, 40}
	println(search(s, 5))  // i = 0
	println(search(s, 15)) // i = 1
	println(search(s, 25)) // i = 2
	println(search(s, 35)) // i = 3
	println(search(s, 45)) // i = 4
}

func search(s []int, v int) int {
	return sort.Search(len(s), func(i int) bool {
		return v < s[i]
	})
}
