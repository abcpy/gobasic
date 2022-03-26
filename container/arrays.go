package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray3(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)
	//arr1 [0 0 0 0 0]

	for i := range arr3 {
		fmt.Printf("%d:%d\n", i, arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr1)
	printArray(arr3)
	//printArray(arr2)

	println("origin arr1 and arr3")
	fmt.Println(arr1)
	fmt.Println(arr3) // [0 0 0 0 0] [2 4 6 8 10]

	println("pointer arr")
	printArray2(&arr1)
	printArray2(&arr3)

	println("origin arr1 and arr3")
	fmt.Println(arr1)
	fmt.Println(arr3) //[100 0 0 0 0]  [100 4 6 8 10]

	println("slice arr")
	printArray3(arr1[:])
	printArray3(arr3[:])

	println("origin arr1 and arr3")
	fmt.Println(arr1)
	fmt.Println(arr3) //[100 0 0 0 0]  [100 4 6 8 10]

}
