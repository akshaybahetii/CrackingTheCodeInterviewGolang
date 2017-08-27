package chap4

func ValidateBSTAndGetMax(t *TreeNode) (bool, int) {
	if t == nil {
		return true, 0
	}
	if t.left == nil && t.right == nil {
		return true, t.data
	}
	resl, l := ValidateBSTAndGetMax(t.left)
	resr, r := ValidateBSTAndGetMax(t.right)

	if resl == false || resr == false {
		return false, -1
	}
	if (l > t.data && t.left != nil) || (r < t.data && t.right != nil) {
		return false, -1
	}

	if t.right != nil {
		return true, r
	}
	return true, l
}

func ValidateBST(t *TreeNode) bool {
	res, _ := ValidateBSTAndGetMax(t)
	return res
}
