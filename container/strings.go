package main

import (
	"fmt"
)

func main() {
	s := "刘健"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()       //E5 88 98 E5 81 A5
	fmt.Println(len(s)) //6

	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	} //(0 刘) (1 健)

}
