/*
	二叉树的中序遍历
	解题思路： 二叉树的遍历天然适合使用递归，前序遍历在遍历左子节点之前记录当前节点数据，
		中序遍历在遍历左子节点之后、右子节点之前记录当前节点数据，后序遍历在遍历右子节点之后记录当前节点数据
*/
package _4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	//记录左子树的遍历结果
	result := inorderTraversal(root.Left)
	//记录当前节点数据
	result = append(result, root.Val)
	//记录右子树的遍历结果
	result = append(result, inorderTraversal(root.Right)...)

	return result
}
