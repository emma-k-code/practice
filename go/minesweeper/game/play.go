package game

import (
	"strconv"
)

// ClickResult 點擊結果
type ClickResult struct {
	Result string `json:"result"`
	Open   []Grip `json:"open"`
}

// Grip 地圖格子內容
type Grip struct {
	Self    string `json:"self"`
	Content string `json:"content"`
	Status  int    `json:"status"` // 1-M 2-Number 3-Zero 4-Boom
}

// 地圖的長寬
var height, width int

// gameMap 遊戲地圖
var gameMap = make(map[int]map[int]int)

// isOpenGrip 已開啟的位置
var isOpenGrip = []string{}

// flagGrip 插旗的位置
var flagGrip = make(map[string]int)

// Init 遊戲地圖初始化
func Init(row, column, m int) {
	height, width = row, column
	gameMap = CreateMap(row, column, m)
	isOpenGrip = []string{}
	flagGrip = make(map[string]int)
}

// CheckClick 檢查點擊結果
func CheckClick(h, w int) ClickResult {
	res := ClickResult{}

	clickGrip := Grip{Self: strconv.Itoa(h) + "_" + strconv.Itoa(w)}

	// 已開啟的位置不可重複執行
	if isOpen(clickGrip.Self) {
		return res
	}

	// 加入已開啟位置資料
	isOpenGrip = append(isOpenGrip, clickGrip.Self)

	// 點擊位置內容
	val := gameMap[h][w]

	if val > 0 {
		res.Result = "save"
		clickGrip.Content = strconv.Itoa(val)
		clickGrip.Status = 2

		res.Open = append(res.Open, clickGrip)
	}

	if val == 0 {
		res.Result = "save"
		clickGrip.Content = ""
		clickGrip.Status = 3

		res.Open = append(res.Open, clickGrip)

		checkAround(h, w, &res)
	}

	if val == -1 {
		// 踩到地雷 game over
		res.Result = "over"
		clickGrip.Content = ""
		clickGrip.Status = 4

		res.Open = append(res.Open, clickGrip)
		gameOver(&res)
	}

	if val >= 0 && isGameClear() {
		res.Result = "clear"
	}

	return res
}

// Flag 格子插旗
func Flag(h, w int) {
	point := strconv.Itoa(h) + "_" + strconv.Itoa(w)

	if _, has := flagGrip[point]; has {
		delete(flagGrip, point)
	} else {
		flagGrip[point] = 1
	}
}

// CheckAroundFlag 檢查四周旗子數量 開啟四周的格子
func CheckAroundFlag(h, w int) ClickResult {
	res := ClickResult{Result: "save"}

	// 點擊位置內容
	clickContent := gameMap[h][w]
	// 插旗數量
	flagCount := 0

	// 八個方位
	around := [8][2]int{
		[2]int{h - 1, w}, [2]int{h - 1, w - 1}, [2]int{h - 1, w + 1},
		[2]int{h, w - 1}, [2]int{h, w + 1},
		[2]int{h + 1, w}, [2]int{h + 1, w - 1}, [2]int{h + 1, w + 1},
	}

	// 計算四周的插旗數量
	for _, aroundPoint := range around {
		checkY := aroundPoint[0]
		checkX := aroundPoint[1]
		if _, has := gameMap[checkY][checkX]; has {
			checkPoint := strconv.Itoa(checkY) + "_" + strconv.Itoa(checkX)
			if isFlag(checkPoint) {
				flagCount++
			}
		}
	}

	// 插旗數量正確 直接開啟四周的格子
	if flagCount == clickContent {
		for _, aroundPoint := range around {
			checkY := aroundPoint[0]
			checkX := aroundPoint[1]
			if val, has := gameMap[checkY][checkX]; has {
				checkPoint := strconv.Itoa(checkY) + "_" + strconv.Itoa(checkX)
				if !isOpen(checkPoint) && !isFlag(checkPoint) {
					checkGrip := Grip{Self: checkPoint}

					// 加入已開啟位置資料
					isOpenGrip = append(isOpenGrip, checkGrip.Self)

					if val == 0 {
						checkGrip.Content = ""
						checkGrip.Status = 3
					}
					if val > 0 {
						checkGrip.Content = strconv.Itoa(val)
						checkGrip.Status = 2
					}
					if val == -1 {
						// 踩到地雷 game over
						res.Result = "over"
						checkGrip.Content = ""
						checkGrip.Status = 4

						res.Open = append(res.Open, checkGrip)
						// 遊戲結束 直接離開
						gameOver(&res)
						break
					}

					// 此次點擊要開啟的格子資料
					res.Open = append(res.Open, checkGrip)

					// 格子為空需要再檢查周圍個格子
					if val == 0 {
						checkAround(checkY, checkX, &res)
					}
				}
			}
		}

		if res.Result != "over" && isGameClear() {
			res.Result = "clear"
		}
	}

	return res
}

// checkAround 檢查四周格子內容，將空格全部開啟
func checkAround(h, w int, res *ClickResult) {
	// 八個方位
	around := [8][2]int{
		[2]int{h - 1, w}, [2]int{h - 1, w - 1}, [2]int{h - 1, w + 1},
		[2]int{h, w - 1}, [2]int{h, w + 1},
		[2]int{h + 1, w}, [2]int{h + 1, w - 1}, [2]int{h + 1, w + 1},
	}

	for _, aroundPoint := range around {
		checkY := aroundPoint[0]
		checkX := aroundPoint[1]
		if val, has := gameMap[checkY][checkX]; has {
			checkPoint := strconv.Itoa(checkY) + "_" + strconv.Itoa(checkX)

			// 格子尚未開啟 or 插旗
			if !isOpen(checkPoint) && !isFlag(checkPoint) {
				checkGrip := Grip{Self: checkPoint}

				if val == 0 {
					checkGrip.Content = ""
					checkGrip.Status = 3
				}
				if val > 0 {
					checkGrip.Content = strconv.Itoa(val)
					checkGrip.Status = 2
				}

				// 此次點擊要開啟的格子資料
				res.Open = append(res.Open, checkGrip)

				// 加入已開啟位置資料
				isOpenGrip = append(isOpenGrip, checkGrip.Self)

				// 格子為空需要再檢查周圍個格子
				if val == 0 {
					checkAround(checkY, checkX, res)
				}
			}
		}
	}
}

// isOpen 檢查該位置是否已開啟
func isOpen(yx string) bool {
	for _, open := range isOpenGrip {
		if open == yx {
			return true
		}
	}
	return false
}

// isFlag 檢查該位置是否已插旗
func isFlag(yx string) bool {
	_, has := flagGrip[yx]
	return has
}

// gameOver 遊戲結束開啟全部格子
func gameOver(res *ClickResult) {
	for h, column := range gameMap {
		for w, val := range column {
			checkPoint := strconv.Itoa(h) + "_" + strconv.Itoa(w)

			// 格子尚未開啟過
			if !isOpen(checkPoint) {
				checkGrip := Grip{Self: checkPoint}

				if val == 0 {
					checkGrip.Content = ""
					checkGrip.Status = 3
				}
				if val > 0 {
					checkGrip.Content = strconv.Itoa(val)
					checkGrip.Status = 2
				}
				if val < 0 {
					checkGrip.Content = ""
					checkGrip.Status = 1
				}

				// 此次點擊要開啟的格子資料
				res.Open = append(res.Open, checkGrip)

				// 加入已開啟位置資料
				isOpenGrip = append(isOpenGrip, checkGrip.Self)
			}
		}
	}
}

// 檢查是否已經過關
func isGameClear() bool {
	total := height * width

	openedCount := len(isOpenGrip) + len(flagGrip)

	return total == openedCount
}
