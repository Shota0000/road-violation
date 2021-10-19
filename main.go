package main

import (
	"road_violation/route"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var (
	ipA        = "192.168.0.3" //userAがいるエッジサーバのIPアドレス
	ipB        = "192.168.0.4" //userBがいるエッジサーバのIPアドレス
	localIp    = "127.0.0.1"
	mongoPortA = "27017"
)

var Session *mgo.Session

func main() {
	dbConnect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	// これなら渡る
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", gin.H{"data": "maoaoa"})
	// })
	route.DefineRoutes(r, Session)
	r.Run()
}

//mongoDB接続
func dbConnect() {
	//docker内部用
	// env := new(on.Local)
	// err := env.Init()
	session, err := mgo.Dial(localIp + ":" + mongoPortA)
	Session = session
	if err != nil {
		panic(err)
	}

}
