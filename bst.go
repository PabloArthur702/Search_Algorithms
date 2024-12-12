package main

import (
	"fmt"
)


type BSTNode struct{
	left *BSTNode
	val int
	right *BSTNode
}

func createNode (val int) *BSTNode{
	return &BSTNode{val:val}
}

func (node *BSTNode) Add(val int){
	if val < node.val {
		if node.left != nil{
			node.left.Add(val)
		}else{
			node.left = createNode(val)
		}
	}else{
		if node.right != nil{
			node.right.Add(val)
		}else{
			node.right = createNode(val)
		}
	}
}

func (node *BSTNode) Search(val int) bool{
	if val == node.val{
		return true
	}else if val < node.val {
		if node.left != nil{
			return node.left.Search(val)
		}
	}else if val > node.val {
		if node.right != nil{
			return node.right.Search(val)
		}
	}
	return false
}

func (node *BSTNode) Min() int{
	if node.left != nil{
		return node.left.Min()
	}else{
		return node.val
	}
}

func (node *BSTNode) Max() int{
	if node.right != nil{
		return node.right.Max()
	}else{
		return node.val
	}
}

func (node *BSTNode) Height() int{
	hLeft := 0
	hRight := 0

	if node.left == nil && node.right == nil{
		return 0
	}
	if node.left != nil{
		hLeft = node.left.Height()
	}
	if node.right != nil{
		hRight = node.right.Height()
	}

	hMax := 0
	if hLeft > hRight{
		hMax = hLeft
	}else{
		hMax = hRight
	}
	
	return hMax + 1
}

func (node *BSTNode) InOrderNav(){
	if node.left != nil{
		node.left.InOrderNav()	
	}
	fmt.Println(node.val)
	if node.right != nil{
		node.right.InOrderNav()
	}
	
}

func main(){
	bst := createNode(10)
	bst.Add(5)
	bst.Add(20)
	bst.Add(15)
	bst.Add(16)
	fmt.Println(bst.Height())
}