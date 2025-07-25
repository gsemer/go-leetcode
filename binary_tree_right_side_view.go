package goleetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	nodes := []int{}
	if root == nil {
		return nodes
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		level := queue.Len()

		for i := 0; i < level; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if i == level-1 {
				nodes = append(nodes, node.Val)
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

	}

	return nodes
}
