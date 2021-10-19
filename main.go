package main

import (
	"road_violation/on"

	"github.com/gin-gonic/gin"
)

var (
	ipA = "192.168.0.3" //userAがいるエッジサーバのIPアドレス
	ipB = "192.168.0.4" //userBがいるエッジサーバのIPアドレス
)

func main() {
	dbConnect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	// route.DefineRoutes(r, session)
	r.Run()
}

//mongoDB接続
func dbConnect() {
	env := new(on.Local)
	env.Init()
	session := env.DB.Session
	print(session)
}
