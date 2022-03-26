package main

import (
	"fmt"
	"learngo/retiver/mock"
	"learngo/retiver/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{
		"this is a fake",
	}

	fmt.Printf("%T %v\n", r, r)
	//mock.Retriever {this is a fake}
	r = &real.Retriever{}
	fmt.Printf("%T %v\n", r, r)
	//*real.Retriever &{ 0s}

	//fmt.Println(download(r))
}
