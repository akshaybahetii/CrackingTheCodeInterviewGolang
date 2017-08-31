package chap8

type Position struct {
	x int
	y int
}

type SQ struct {
	allowed bool
}

func GetPath(b [][]SQ, start Position, path *[]Position) {
	if start.x+1 < len(b) && !b[start.x+1][start.y].allowed {
		p := Position{start.x + 1, start.y}
		*path = append(*path, p)
		GetPath(b, p, path)
} else if start.y+1 < len(b[0]) && !b[start.x][start.y+1].allowed {
		p := Position{start.x, start.y + 1}
		*path = append(*path, p)
		GetPath(b, p, path)
	}
}
