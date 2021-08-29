/*
	200. 岛屿数量
	解题思路：使用dfs搜索陆地，并对陆地染色，完成一次dfs就代表一个岛屿
*/

package _00

func numIslands(grid [][]byte) int {
	var isLandNum int

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 || grid[x][y] != '1' {
			return
		}
		grid[x][y]++
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x, y+1)

		return
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				isLandNum++
			}
		}
	}

	return isLandNum
}
