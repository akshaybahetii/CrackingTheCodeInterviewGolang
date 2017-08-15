package chap1

func CheckStringRotation(s string, r string) bool {
	if len(s) != len(r) {
		return false
	}
	ss := []rune(s)
	rr := []rune(r)
	startfound := false
	start := -1
	i := 0
	j := 0
	for i < len(ss) && j < len(rr) {
		x := ss[i]
		y := rr[j]
		if startfound && x != y {
			//The start we found was wrong restart
			startfound = false
			i = 0
		}
		if x == y {
			//might be start of of string s
			if !startfound {
				startfound = true
				start = j
				i++
			} else {
				i++
			}
		}
		j++
	}
	//Check if 0:start is a substring
	if startfound && string(ss[(len(ss)-start):]) == string(rr[:(start)]) {
		return true
	}
	return false
}
