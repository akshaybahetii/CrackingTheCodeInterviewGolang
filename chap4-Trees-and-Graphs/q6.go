package chap4

func FindNodeInBST(t *TreeNode, x int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.data == x {
		//Find the next inorder node.
		if t.left != nil {
			return t.left
		} else if t.right != nil {
			return t.right
		} else {
			if t.parent == nil {
				return t
			}
			if t.parent.parent == nil {
				return nil
			}
			return t.parent.parent.right
		}
	}
	if t.data < x {
		return FindNodeInBST(t.right, x)
	} else {
		return FindNodeInBST(t.left, x)
	}
}
