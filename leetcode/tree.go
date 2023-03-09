package leetcode

import "fmt"

/**
 * Definition for binary tree node
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(nums []any) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root} // 节点指针队列
	index := 0                 // 结点指针队列索引，记录遍历数组时当前节点的父节点
	isLeft := true

	for i := 1; i < len(nums); i++ {
		var node *TreeNode = nil
		if nums[i] != nil {
			node = &TreeNode{Val: nums[i].(int)}
		}
		parent := queue[index]
		if isLeft {
			parent.Left = node
		} else {
			parent.Right = node
			index++
		}
		isLeft = !isLeft
		if node != nil {
			queue = append(queue, node)
		}
	}
	return root
}

func (root *TreeNode) Print() {
	str := root.outputString()
	fmt.Printf(str + "\n")
}

func (root *TreeNode) outputString() string {
	if root == nil {
		return fmt.Sprintf("[]")
	}
	str := fmt.Sprintf("[")
	queue := []*TreeNode{root}
	for i := 0; i < len(queue); i++ {
		p := queue[i]
		if i == 0 {
			str = str + fmt.Sprintf("%v", p.Val)
		} else {
			str = str + fmt.Sprintf(", %v", p.Val)
		}
		if p.Left != nil {
			queue = append(queue, p.Left)
		}
		if p.Right != nil {
			queue = append(queue, p.Right)
		}
	}
	str = str + fmt.Sprintf("]")
	return str
}
