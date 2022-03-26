package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	//fmt.Println(root)  //{0 <nil> <nil>}
	root = tree.Node{Value: 3}
	//fmt.Println(root) //{3 <nil> <nil>}
	root.Left = &tree.Node{}
	//fmt.Println(root) //{3 0xc000096060 <nil>}
	root.Right = &tree.Node{5, nil, nil}
	//fmt.Println(root) //{3 0xc000096060 0xc000096078}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	/*	fmt.Println(root.Left)
		fmt.Println()
		root.Print() //3*/
	root.Right.Left.SetValue(300)
	//root.Right.Left.Print()

	fmt.Println("middle traverse")
	root.Traverse()

	/*fmt.Println()
	node := myTreeNode{&root}
	node.postOrder()*/

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount += 1
	})

	fmt.Println("Node count:", nodeCount)

	/*var pRoot *treeNode
	pRoot.setValue(500)
	pRoot = &root
	pRoot.setValue(800)
	pRoot.print()*/

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, nil},
	//}
	//fmt.Println(nodes)

}
