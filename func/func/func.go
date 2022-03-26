package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}

	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

//func div2(a, b int) (q, r int) {
//	q = a / b
//	r = a % b
//	return q, r
//}

func main() {
	fmt.Println(eval(3, 4, "x"))
	fmt.Println(div(13, 3))
	//div2(13, 3)

	result := apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))

	}, 3, 4)

	fmt.Println(result)
	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b) //4 3

	c, d := swap2(a, b)
	fmt.Println(c, d) //3 4 
}
