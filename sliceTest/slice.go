package main

import "fmt"

func main() {
	//merge1()
	//merge2()
	//merge3()
	//merge4()
	compare()
}

// 先改s2 对应的底层数据，再改了s3的，所以是对的
func merge1()  {
	s1 := []int{0,1,2,5,6,7}
	s2 := []int{3,4}
	s1 = append(s1[:3],append(s2, s1[3:]...)...)
	fmt.Printf("merge1 is: (%+v)", s1)
}

// 先改s3 对应的底层数据，再往把s3截取来里面，所以是错的
func merge2()  {
	s1 := []int{0,1,2,5,6,7}
	s2 := []int{3,4}
	s1 = append(append(s1[:3],s2...), s1[3:]...)
	fmt.Printf("merge2 is: (%+v)", s1)
}

func merge3()  {
	s1 := []int{0,1,2,5,6,7}
	s2 := []int{3,4}
	s3 := append(s2, s1[3:]...)
	fmt.Printf("merge3 s1 (%+v), s2(%+v), s3 (%+v)", s1, s2, s3)
	s1 = append(s1[:3],s3...)
	fmt.Printf("merge3 s1 (%+v), s2(%+v), s3 (%+v)", s1, s2, s3)
}

func merge4()  {
	s1 := []int{0,1,2,5,6,7}
	s2 := []int{3,4}
	fmt.Printf("merge4 原始的 s1 (%p), s2(%p), cap s1(%d)\n", s1, s2, cap(s1))
	s3 := append(s1[:3], s2...)
	fmt.Printf("merge4 s1 (%p), s2(%p), s3(%p)\n", s1, s2, s3)
	fmt.Printf("merge4 s1 (%+v), s2(%+v)\n", s1, s2)
	//fmt.Printf("s1 ptr == s2 ptr (%t)", )
	//s1 = append(s3, s1[3:]...)
	//fmt.Printf("merge4 s1 (%p), s2(%p), s3(%p)\n", s1, s2, s3)
	//fmt.Printf("merge4 s1 (%+v), s2(%+v), s3 (%+v)\n", s1, s2, s3)
}

func compare()  {
	s1 := []int{0}
	s2 := []int{1}
	s1 = append(s2[:1], s1...)
	fmt.Printf("s1 (%+v)", s1)
}