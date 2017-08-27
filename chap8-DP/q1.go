package chap8

func GetWaysStairs(n int, res *map[int]int) int {
	if res == nil {
		x := make(map[int]int)
		res = &x
	}
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		m := *res
		v, ok := m[n]
		if ok {
			return v
		} else {
			x := GetWaysStairs(n-1, res) + GetWaysStairs(n-3, res) + GetWaysStairs(n-2, res) + 3
			m[n] = x
			return x
		}
	}
}
