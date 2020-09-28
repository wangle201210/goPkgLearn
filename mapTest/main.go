package main

import "fmt"

func main() {
	mf := make(map[int]int)
	fmt.Println(len(mf))
	uptest(mf)
	fmt.Println("after uptest ", len(mf))
	fmt.Println("done")
}

func uptest(m map[int]int) {
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
	fmt.Println("in uptest ", len(m))
}