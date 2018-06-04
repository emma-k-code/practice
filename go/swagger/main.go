// Package main Swagger 測試用 API
//
// 此為說明文字
//
//     Schemes: http
//     BasePath: /api
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type GetUserParam struct {
	// 會員ID
	Id int `json:"id"`
}

type ResponseData struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Ret  GetUserRet `json:"ret"`
}

type GetUserRet struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(c echo.Context) error {
	// swagger:operation GET /user User getSingleUser
	//
	// get a user by userID
	//
	// This will show a user info
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: query
	//   description: 會員ID
	//   required: true
	//   type: string
	// responses:
	//   200:
	//     description: "成功回傳資料"
	//     schema:
	//       type: object
	//       required:
	//       - code
	//       - msg
	//       - ret
	//       properties:
	//         code:
	//           description: 錯誤代碼
	//           type: integer
	//         msg:
	//           description: 錯誤訊息
	//           type: string
	//         ret:
	//           description: 回傳資料
	//           type: object
	//           properties:
	//             id:
	//               description: 會員ID
	//               type: integer
	//             name:
	//               description: 會員名稱
	//               type: string
	//             email:
	//               description: 會員Email
	//               type: string
	//       example:
	//         code: 0
	//         msg: 測試用API
	//         ret:
	//           id: 0
	//           name: test
	//           email: test@mail.com
	id, _ := strconv.Atoi(c.QueryParam("id"))

	res := ResponseData{
		Code: 0,
		Msg:  "測試用API",
		Ret:  GetUserRet{Id: id, Name: "test", Email: "test@gmail.com"},
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	// c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	return c.JSON(http.StatusOK, res)
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 靜態檔
	e.Static("/swagger", "html")
	e.Static("/swagger/json", "json")

	// test api
	e.GET("/api/user", getUser)

	e.Logger.Fatal(e.Start(":8081"))
}
