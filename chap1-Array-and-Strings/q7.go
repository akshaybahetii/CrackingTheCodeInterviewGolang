package chap1

func RotateMatrix(mat [3][3]int) [3][3]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if j < i {
				t := mat[i][j]
				mat[i][j] = mat[j][i]
				mat[j][i] = t
			}
		}
	}
	return mat
}
