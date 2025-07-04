package gameoflife

func gameOfLife(board [][]int) {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])

	b := make([][]int, n)
	for i := range board {
		b[i] = make([]int, m)
		copy(b[i], board[i])
	}

	dx := []int{1, 0, -1, 0, 1, 1, -1, -1}
	dy := []int{0, 1, 0, -1, 1, -1, 1, -1}

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			liveNeighbors := 0
			for k := 0; k < 8; k++ {
				nx := x + dx[k]
				ny := y + dy[k]

				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					liveNeighbors += b[ny][nx]
				}
			}

			if b[y][x] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					board[y][x] = 0
				}
			} else {
				if liveNeighbors == 3 {
					board[y][x] = 1
				}
			}
		}
	}
}