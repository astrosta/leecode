/*
	529. 扫雷游戏
	解题思路： 对每个节点广度优先遍历，若发现地雷停止遍历
*/

package _29

func updateBoard(board [][]byte, click []int) [][]byte {
	var dirX = [8]int{-1, 0, 1, 1, 1, 0, -1, -1}
	var dirY = [8]int{1, 1, 1, 0, -1, -1, -1, 0}
	x := click[0]
	y := click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	var bfs func(x, y int)
	bfs = func(x, y int) {
		var cnt byte = '0'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) ||
				ty < 0 || ty >= len(board[0]) {
				continue
			}
			if board[tx][ty] == 'M' {
				cnt = cnt + 1
			}
		}

		if cnt != '0' {
			board[x][y] = cnt
		} else {
			board[x][y] = 'B'
			for i := 0; i < 8; i++ {
				tx, ty := x+dirX[i], y+dirY[i]
				if tx < 0 || tx >= len(board) ||
					ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
					continue
				}
				bfs(tx, ty)
			}
		}
	}

	bfs(x, y)

	return board
}
