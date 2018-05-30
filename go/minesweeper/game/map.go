package game

import (
	"math/rand"
	"strconv"
	"time"
)

// CreateMap 建立遊戲地圖
func CreateMap(row, column, m int) map[int]map[int]int {
	// 空白的遊戲地圖
	gameMap := make(map[int]map[int]int, row)
	for height := 0; height < row; height++ {
		gameMap[height] = make(map[int]int, column)
		for width := 0; width < column; width++ {
			gameMap[height][width] = 0
		}
	}

	// 亂數決定地雷位置
	mineIndex := GetMineIndex(m, row, column)

	// 設定地雷地圖位置
	for _, index := range mineIndex {
		// 0:高 1:寬
		gameMap[index[0]][index[1]] = -1
	}

	// 計算每個格子周圍的地雷數
	for h, columnMap := range gameMap {
		for w, val := range columnMap {
			// 沒有地雷需要計算八個方向的格子地雷數
			if val == 0 {
				// 上
				if val, has := gameMap[h-1][w]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}
				// 下
				if val, has := gameMap[h+1][w]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}

				// 左測上~下
				if val, has := gameMap[h-1][w-1]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}
				if val, has := gameMap[h][w-1]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}
				if val, has := gameMap[h+1][w-1]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}

				// 右側上~下
				if val, has := gameMap[h-1][w+1]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}
				if val, has := gameMap[h][w+1]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}
				if val, has := gameMap[h+1][w+1]; has {
					if val == -1 {
						gameMap[h][w]++
					}
				}
			}
		}
	}

	return gameMap
}

// GetMineIndex 取得地雷位置
func GetMineIndex(mineCount, row, column int) map[int][2]int {
	// 亂數最大值
	max := row * column

	// 地雷位置
	mineIndex := make(map[int][2]int)

	for i := 0; i < mineCount; i++ {
		index := MineRand(max)

		// 高 and 寬的位置
		mineIndex[i] = [2]int{index / column, index % column}
	}

	return mineIndex
}

// MineRand 亂數產生器
func MineRand(max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max)
}

// BlankMapHTML 輸出空白的地圖Html
func BlankMapHTML(row, column int) string {
	gameHTML := ""

	for h := 0; h < row; h++ {
		gameHTML += `<tr height="40px" align="center">`
		for w := 0; w < column; w++ {
			num := strconv.Itoa(h) + "_" + strconv.Itoa(w)
			gameHTML += `<td width="40px" id="` + num + `">`
			gameHTML += `<span id="content" style="display:none;"></span>`
			gameHTML += `<span id="icon" style="display:none;"></span>`
			gameHTML += `<img id="imgM" src="icon/bomb.png" height="36 width="36" style="display:none;">`
			gameHTML += `<span id="flag" class="glyphicon glyphicon-flag" style="display:none"></span>`
			gameHTML += `</td>`
		}
		gameHTML += `</tr>`
	}
	return gameHTML
}
