package node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SortedArrayToBST 108. Convert Sorted Array to Binary Search Tree(有排過，由小到大)
func SortedArrayToBST(nums []int) *TreeNode {
	//先算長度
	n := len(nums)

	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	//因二元樹要平衡，root會是陣列中最中間的元素
	m := n / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  SortedArrayToBST(nums[:m]),
		Right: SortedArrayToBST(nums[m+1:]),
	}
}
