package main

import "fmt"

func printSlice(slice []int) {
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))
}

func main() {
	var s []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	fmt.Println("Crating slice")
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("len=%d, cap=%d\n", len(s3), cap(s3))

	/*len=4, cap=4
	len=16, cap=16
	len=10, cap=32*/

	fmt.Println("Coping slice")
	copy(s2, s1)
	fmt.Println(s2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2) //[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	fmt.Println(s2)               //[4 6 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(s2), cap(s2)) //14 15

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	fmt.Println(s2)               //4 6 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(s2), cap(s2)) //13 15

}
