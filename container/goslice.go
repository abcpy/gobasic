package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Println(s1, s2)
	// s1: [3 4 5 6]
	// s2: [6 7]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	//s1=[3 4 5 6], len(s1)=4, cap(s1)=5
	//s2=[6 7], len(s2)=2, cap(s2)=2

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3, s4, s5=", s3, s4, s5)
	//s3, s4, s5= [6 7 10] [6 7 10 11] [6 7 10 11 12]
	//s4 and s5 no longer view arr
	fmt.Println("arr", arr)
	//arr [0 1 3 4 5 6 7]
}
