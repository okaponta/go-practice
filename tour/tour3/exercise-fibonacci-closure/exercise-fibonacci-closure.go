package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	isFirst, before, now := true, 1, 0
	return func() int {
		if isFirst {
			isFirst = false
			return 0
		}
		before, now = now, before+now
		return now
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
