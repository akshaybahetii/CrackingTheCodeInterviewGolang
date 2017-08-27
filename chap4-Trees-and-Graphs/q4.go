package chap4

func GetHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}
	l := GetHeight(t.left)
	r := GetHeight(t.right)

	if l > r {
		return l + 1
	}
	return r + 1

}

func CheckBalanced(t *TreeNode) bool {
	l := GetHeight(t.left)
	r := GetHeight(t.right)
	if l == r+1 || r == l+1 || l == r {
		return true
	}
	return false
}
