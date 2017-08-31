package chap8

// Check if t1 and t2 are empty return
//i

func Move(s *[]int, d *[]int) {
	x := (*s)[0]
	*s = (*s)[1:]
	*d = append([]int{x}, (*d)[0:]...)
}

func TOI(n int, t1 *[]int, t2 *[]int, t3 *[]int) {
	if n <= 0 {
		return
	}

	TOI(n-1, t1, t3, t2)
	Move(t1, t3)
	TOI(n-1, t2, t1, t3)
}
