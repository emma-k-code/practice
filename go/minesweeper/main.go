package main

import (
	"fmt"
	"net/http"
	"strconv"

	Game "./game"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createGame(c echo.Context) error {
	// 取得傳入的 param
	row, _ := strconv.Atoi(c.QueryParam("row"))
	column, _ := strconv.Atoi(c.QueryParam("column"))
	m, _ := strconv.Atoi(c.QueryParam("m"))

	// 取得遊戲地圖
	gameMapHtml := Game.CreateMap(row, column, m)

	return c.String(http.StatusOK, gameMapHtml)

	defer func() {
		if p := recover(); p != nil {
			fmt.Sprintf("cliententer panic: %v", p)
		}
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

	e.Logger.Fatal(e.Start(":8080"))
}
