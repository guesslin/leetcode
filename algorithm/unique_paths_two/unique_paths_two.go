package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	steps := make([][]int, m)
	steps[0] = append(steps[0], 1)
	for col := 1; col < n; col++ {
		if obstacleGrid[0][col] != 1 {
			steps[0] = append(steps[0], steps[0][col-1])
		} else {
			steps[0] = append(steps[0], 0)
		}
	}
	for row := 1; row < m; row++ {
		if obstacleGrid[row][0] != 1 {
			steps[row] = append(steps[row], steps[row-1][0])
		} else {
			steps[row] = append(steps[row], 0)
		}
		for col := 1; col < n; col++ {
			if obstacleGrid[row][col] != 1 {
				steps[row] = append(steps[row], steps[row][col-1]+steps[row-1][col])
			} else {
				steps[row] = append(steps[row], 0)
			}
		}
	}
	return steps[m-1][n-1]
}
