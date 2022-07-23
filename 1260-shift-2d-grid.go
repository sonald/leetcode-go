package main

func shiftGrid(grid [][]int, k int) [][]int {
	R, C := len(grid), len(grid[0])
	k = k % (R * C)

	for i := 0; i < k; i++ {
		last := grid[R-1][C-1]
		for r := R - 1; r >= 0; r-- {
			for c := C - 1; c >= 0; c-- {
				prevc, prevr := c-1, r
				if c == 0 {
					prevc, prevr = C-1, (r-1+R)%R
				}
				grid[r][c] = grid[prevr][prevc]
			}
		}

		grid[0][0] = last
	}

	return grid
}

func shiftGrid2(grid [][]int, k int) [][]int {
	if k == 0 {
		return grid
	}
	R, C := len(grid), len(grid[0])
	k = k % (R * C)
	flat := make([]int, R*C+k)
	for i := 0; i < R; i++ {
		copy(flat[i*C:i*C+C], grid[i])
	}

	copy(flat[k:], flat[:R*C])
	copy(flat[:k], flat[R*C:])

	for i := 0; i < R; i++ {
		copy(grid[i], flat[i*C:i*C+C])
	}

	return grid
}
