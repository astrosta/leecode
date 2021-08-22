/*
	从前序与中序遍历序列构造二叉树
	解题思路：转换成找根节点的值和左子树、右子树对应的序列化数据，这样就能使用递归构造树的每个节点
*/

package _05

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := new(TreeNode)
	node.Val = preorder[0]

	var i int
	for i = 0; i < len(preorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return node
}
