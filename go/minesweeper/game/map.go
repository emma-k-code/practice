package game

import (
	"math/rand"
	"strconv"
	"time"
)

// CreateMap 建立遊戲地圖
func CreateMap(row, column, m int) [][]int {
	// 空白的遊戲地圖
	gameMap := make([][]int, row)
	for height := 0; height < row; height++ {
		gameMap[height] = make([]int, column)
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
				// 八個方位
				around := [8][2]int{
					[2]int{h - 1, w}, [2]int{h - 1, w - 1}, [2]int{h - 1, w + 1},
					[2]int{h, w - 1}, [2]int{h, w + 1},
					[2]int{h + 1, w}, [2]int{h + 1, w - 1}, [2]int{h + 1, w + 1},
				}

				for _, aroundPoint := range around {
					checkY := aroundPoint[0]
					checkX := aroundPoint[1]
					if checkY >= 0 && checkY < row && checkX >= 0 && checkX < column {
						if gameMap[checkY][checkX] == -1 {
							gameMap[h][w]++
						}
					}
				}
			}
		}
	}

	return gameMap
}

// GetMineIndex 取得地雷位置
func GetMineIndex(mineCount, row, column int) [][2]int {
	// 亂數最大值
	max := row * column

	// 亂數表 (避免出現重複亂數)
	randNum := make([]int, max)
	for i := 0; i < max; i++ {
		randNum[i] = i
	}

	// 地雷位置
	mineIndex := make([][2]int, mineCount)

	for i := 0; i < mineCount; i++ {
		// 取除亂數表中的亂數
		index := MineRand(len(randNum) - 1)

		// 從亂數表中移除已出現的數字
		randNum = append(randNum[:index], randNum[index+1:]...)

		// 高 and 寬的位置
		mineIndex[i] = [2]int{randNum[index] / column, randNum[index] % column}
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
