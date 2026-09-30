package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		prev := queue[0].Val

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				if curr.Val%2 == 0 || (i > 0 && curr.Val <= prev) {
					return false
				}
			} else {
				if curr.Val%2 != 0 || (i > 0 && curr.Val >= prev) {
					return false
				}
			}

			prev = curr.Val

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		level++
	}

	return true
}

// Helper function to build a tree from a level-order slice containing nil for null nodes
func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		curr := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != nil {
			curr.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, curr.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			curr.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}

func intPtr(v int) *int {
	return &v
}

func main() {
	vals := []*int{
		intPtr(1), intPtr(10), intPtr(4),
		intPtr(3), nil, intPtr(7), intPtr(9),
		intPtr(12), intPtr(8), intPtr(6), nil, nil, intPtr(2),
	}

	root := buildTree(vals)
	result := isEvenOddTree(root)
	
	fmt.Println("Is Even-Odd Tree:", result)
}