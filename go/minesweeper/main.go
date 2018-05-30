package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	Game "./game"
)

type reqParam struct {
	Name  string `json:"name"`
	Param string `json:"param"`
}

type resData struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func keep(ws *websocket.Conn) bool {
	inputData := reqParam{}
	outputData := resData{}
	//使用迴圈持續接收資料
	for {
		// 讀取ws
		mt, input, err := ws.ReadMessage()
		if err != nil {
			fmt.Sprintf("ws.ReadMessage err: %s\n", err)
			ws.Close()
			return true
		}

		err = json.Unmarshal(input, &inputData)
		if err != nil {
			fmt.Sprintf("json decode err: %s\n", err)
			return true
		}

		outputData.Name = inputData.Name
		outputData.Data = gameSwitch(inputData.Name, inputData.Param)

		output, err := json.Marshal(outputData)
		if err != nil {
			fmt.Sprintf("json decode err: %s\n", err)
			return true
		}

		err = ws.WriteMessage(mt, output)
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

func gameSwitch(name string, param string) interface{} {
	switch name {
	case "create":
		return createGame(param)
	case "click":
		return clickMap(param)
	case "flag":
		return clickFlag(param)
	case "check_around_flag":
		return checkAroundFlag(param)
	}

	return ""
}

func createGame(param string) string {
	params := strings.Split(param, ",")
	row, _ := strconv.Atoi(params[0])
	column, _ := strconv.Atoi(params[1])
	m, _ := strconv.Atoi(params[2])

	Game.Init(row, column, m)

	return Game.BlankMapHTML(row, column)
}

func clickMap(clickPoint string) Game.ClickResult {
	// 點擊位置
	data := strings.Split(clickPoint, "_")
	h, _ := strconv.Atoi(data[0])
	w, _ := strconv.Atoi(data[1])

	return Game.CheckClick(h, w)
}

func clickFlag(param string) string {
	// 點擊位置
	data := strings.Split(param, "_")
	h, _ := strconv.Atoi(data[0])
	w, _ := strconv.Atoi(data[1])

	Game.Flag(h, w)

	return param
}

func checkAroundFlag(clickPoint string) Game.ClickResult {
	// 點擊位置
	data := strings.Split(clickPoint, "_")
	h, _ := strconv.Atoi(data[0])
	w, _ := strconv.Atoi(data[1])

	return Game.CheckAroundFlag(h, w)
}

func websocketStart(c echo.Context) error {
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

	// websocket
	e.GET("/ws", websocketStart)

	e.Logger.Fatal(e.Start(":8080"))
}
