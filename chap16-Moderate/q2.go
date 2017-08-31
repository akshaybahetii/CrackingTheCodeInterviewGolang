package chap16

func GetFreq(bookMap *map[string]int, book []string, word string) int {
	if len(*bookMap) == 0 {
		for _, w := range book {
			(*bookMap)[w]++
		}
	}
	return (*bookMap)[word]
}
