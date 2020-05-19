package main

import (
	"net/http"
	"github.com/labstack/echo"
  	// "github.com/labstack/echo/middleware"
)

func main(){
	e := echo.New()

	//CORSの設定(vueのプロジェクトをGOで立てたlocalサーバーで起動する時は不要)
	// e.Use(middleware.CORS())

	// npm run buildでビルドしたものをgoで起動  corsも不要になる
	//  /でアクセスしたときのルーティング設定
	e.Static("/", "dist/")

	// リクエストに対するHandler
	e.GET("/getTitle", getTitle)
	e.GET("/getName/:name", getName)
	e.POST("/postName", postName)
	e.POST("/postCompany", postCompany)

	// local サーバー
	e.Logger.Fatal(e.Start(":8000"))
}

// GETリクエスト
func getTitle(c echo.Context) error {
	return c.String(http.StatusOK, "New Game")
}

// パラメータ付きのGETリクエスト
func getName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

// application/x-www-form-urlencoded データのPOSTリクエスト
func postName(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}

//JSON受け取り用の構造体
type JsonParam struct {
	Company string `json:"company"`
	Works string `json:"works"`
}

// JSONデータのPOSTリクエスト
func postCompany(c echo.Context) error {
	param := new(JsonParam)
	//バインドしてJSON取得
	if err := c.Bind(param); err != nil {
		return err
	}
	//JSONを返す
	return c.JSON(http.StatusOK, param)
}