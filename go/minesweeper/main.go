package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	Game "./game"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func createGame(c echo.Context) error {
	// 取得傳入的 param
	row, _ := strconv.Atoi(c.QueryParam("row"))
	column, _ := strconv.Atoi(c.QueryParam("column"))
	m, _ := strconv.Atoi(c.QueryParam("m"))

	// 取得遊戲地圖
	gameMapHTML := Game.CreateMap(row, column, m)

	return c.String(http.StatusOK, gameMapHTML)

	defer func() {
		if p := recover(); p != nil {
			fmt.Sprintf("cliententer panic: %v", p)
		}
	}()
	return nil
}

func keep(ws *websocket.Conn) bool {
	//使用迴圈持續接收資料
	for {
		// 讀取ws
		mt, param, err := ws.ReadMessage()
		if err != nil {
			fmt.Sprintf("ws.ReadMessage err: %s\n", err)
			ws.Close()
			return true
		}

		params := strings.Split(string(param), ",")
		row, _ := strconv.Atoi(params[0])
		column, _ := strconv.Atoi(params[1])
		m, _ := strconv.Atoi(params[2])

		// 取得遊戲地圖
		gameMapHTML := Game.CreateMap(row, column, m)

		err = ws.WriteMessage(mt, []byte(gameMapHTML))
		if err != nil {
			fmt.Sprintf("WriteMessage err: %s\n", err)
		}
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Sprintf("cliententer panic: %v", p)
		}
	}()
	return true
}

func gameWebsocket(c echo.Context) error {
	//建立websocket連線
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	//錯誤處理
	if err != nil {
		fmt.Sprintf("upgrader.Upgrade err: %s\n", err)
		return nil
	}
	//開始接收客端資料
	keep(ws)
	defer func() {
		if p := recover(); p != nil {
			fmt.Sprintf("cliententer panic: %v", p)
		}
		ws.Close()
	}()
	return nil
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 靜態檔
	e.Static("/", "index.html")
	e.Static("/js", "js")
	e.Static("/css", "css")
	e.Static("/icon", "icon")

	// 建立遊戲
	e.GET("/game/create", createGame)
	// websocket
	e.GET("/ws", gameWebsocket)

	e.Logger.Fatal(e.Start(":8080"))
}
