package goleetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}

	var buildTree func(start int, end int) *TreeNode
	buildTree = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}

		middle := (start + end) / 2

		return &TreeNode{
			Val:   nums[middle],
			Left:  buildTree(start, middle-1),
			Right: buildTree(middle+1, end),
		}
	}

	return buildTree(0, len(nums)-1)
}
