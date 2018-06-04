go-swagger 範例
=======================
- 使用 echo 建立 API 與 API 文檔頁面
- 使用 go-swagger 產生 swagger 的 json 檔

## 檔案說明
> main.go - API 本體、註解

> swagger/swagger.json - 用 go-swagger 解析註解後得到的 json 檔

> html/* - Swagger API 網頁檔案

## 使用說明
> 建立文檔
> ```go
> $GOPATH/bin/swagger generate spec -o json/swagger.json
> ```

> 啟用 API、API文件網頁
> ```go
> go run main.go
> ```

## 執行後的文檔網址
http://127.0.0.1:8081/swagger