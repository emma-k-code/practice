package game

import (
	"math/rand"
	"strconv"
	"time"
)

func CreateMap(row, column, m int) string {
	// 總格子數
	total := row * column

	// 遊戲地圖
	gameMap := make(map[int]int)

	// 亂數決定地雷位置
	mineIndex := GetMineIndex(m, total)

	// 先將地圖全部設為 0 (無地雷)
	for i := 1; i <= total; i++ {
		gameMap[i] = 0
	}

	// 設定地雷Index
	for _, index := range mineIndex {
		gameMap[index] = -1
	}

	// 計算格子附近的地雷數
	for key, val := range gameMap {
		if val == 0 {
			// 判斷是否為最右側
			if key%row != 0 {
				// 右
				if gameMap[key+1] == -1 {
					gameMap[key] = gameMap[key] + 1
				}
				// 右下
				if _, has := gameMap[key+row+1]; has {
					if gameMap[key+row+1] == -1 {
						gameMap[key] = gameMap[key] + 1
					}
				}
				// 右上
				if _, has := gameMap[key-row-1]; has {
					if gameMap[key-row-1] == -1 {
						gameMap[key] = gameMap[key] + 1
					}
				}
			}

			// 判斷是否為最左側
			if key%row != 1 {
				// 右
				if gameMap[key-1] == -1 {
					gameMap[key] = gameMap[key] + 1
				}
				// 左下
				if _, has := gameMap[key+row-1]; has {
					if gameMap[key+row-1] == -1 {
						gameMap[key] = gameMap[key] + 1
					}
				}
				// 左上
				if _, has := gameMap[key-row+1]; has {
					if gameMap[key-row+1] == -1 {
						gameMap[key] = gameMap[key] + 1
					}
				}
			}

			// 下
			if _, has := gameMap[key+row]; has {
				if gameMap[key+row] == -1 {
					gameMap[key] = gameMap[key] + 1
				}
			}
			// 上
			if _, has := gameMap[key-row]; has {
				if gameMap[key-row] == -1 {
					gameMap[key] = gameMap[key] + 1
				}
			}
		}
	}

	// 轉換為地圖html
	return MapHTML(row, column, m, gameMap)
}

func GetMineIndex(mineCount, max int) map[int]int {
	mineIndex := make(map[int]int)
	for i := 0; i < mineCount; i++ {
		mineIndex[i] = MineRand(max)
	}

	return mineIndex
}

func MineRand(max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	r_num := rand.Intn(max) + 1
	return r_num
}

func MapHTML(row, column, m int, gameMap map[int]int) string {
	gameHtml := ""
	for i := 0; i < column; i++ {
		gameHtml += `<tr height="40px" align="center">`
		for h := 0; h < row; h++ {
			key := i*row + (h + 1)

			gameHtml += `<td width="40px">`

			gameHtml += `<span id="content" style="display:none;">`
			// 該地圖格子的內容
			if gameMap[key] == -1 {
				gameHtml += "M"
			}
			if gameMap[key] > 0 {
				gameHtml += strconv.Itoa(gameMap[key])
			}
			gameHtml += `</span>`

			gameHtml += `<span id="icon" style="display:none;"></span>`

			// 該地圖格子是地雷，加上地雷圖示
			if gameMap[key] == -1 {
				gameHtml += `<img id="imgM" src="icon/bomb.png" height="36 width="36" style="display:none;">`
			}

			gameHtml += `<span id="flag" class="glyphicon glyphicon-flag" style="display:none"></span>`

			gameHtml += `</td>`
		}
		gameHtml += `</tr>`
	}

	return gameHtml
}
