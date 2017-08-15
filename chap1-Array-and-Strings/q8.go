package chap1

func ZeroMatrix(matP *[][]int) *[][]int {
	mat := *matP
	lenOfRow := len(mat)
	lenOfCol := len(mat[0])

	rows := make([]int, lenOfRow)
	colmuns := make([]int, lenOfCol)

	for i := 0; i < lenOfRow; i++ {
		for j := 0; j < lenOfCol; j++ {
			if mat[i][j] == 0 {
				rows[i] = -1
				colmuns[i] = -1
			}
		}
	}

	for i, x := range rows {
		if x == -1 {
			for j := 0; j < lenOfRow; j++ {
				mat[i][j] = 0
			}
		}
	}

	for i, x := range colmuns {
		if x == -1 {
			for j := 0; j < lenOfCol; j++ {
				mat[j][i] = 0
			}
		}
	}
	return &mat
}
