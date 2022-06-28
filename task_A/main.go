package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i < t; i++ {
		var s1, s2 int
		fmt.Scan(&s1, &s2)
		fmt.Println(s1 + s2)
	}
}
