package chap4

func isChild(t *TreeNode) bool {
	if t.left == nil && t.right == nil {
		return true
	}
	return false
}

func hasOneChild(t *TreeNode) bool {
	if (t.right == nil && t.left != nil) || (t.right != nil && t.left == nil) {
		return true
	}
}
func DeleteBST(t *TreeNode, x int) {
	//Delete child node.
	if t.right != nil && t.right.data == x && isChild(t.right) {
		t.right = nil
	}
	if t.left != nil && t.left.data == x && isChild(t.left) {
		t.left = nil
	}

	//Delete parent with one child.
	if t.right != nil && hasOneChild(t.right) && t.right.data == x {
		if t.right.left != nil {
			t.right = t.right.left
		} else {
			t.right = t.right.right
		}
	}
	if t.left != nil && hasOneChild(t.left) && t.left.data == x {
		if t.left.left != nil {
			t.left = t.left.left
		} else {
			t.left = t.left.right
		}
	}

	//Delete parent with 2 children.

}
