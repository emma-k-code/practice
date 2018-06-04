Go 相關練習
=======================

## json_test
各套件的 json 處理速度測試

| 測試套件         |
|-----------------|
| Std Library     |
| JsonIter        |

> 測試指令
> ```go
> go test -bench=. -benchmem
> ```

## minesweeper
將之前寫的 PHP 踩地雷用 Go 改寫
- 用 echo 套件建立連線 (port: 8080)
- 已 websocket 進行遊戲操作

> 使用執行檔
> ```go
> ./minesweeper
> ```

## swagger
使用 go-swagger 建立 API 文檔