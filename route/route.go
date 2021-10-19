package route

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Car struct {
	Id        bson.ObjectId          `bson:"_id"`
	Name      string                 `bson:"name"`
	Home      string                 `bson:"home"`
	violation map[string]interface{} `bson:"violation"`
}

var car []Car

func DefineRoutes(router *gin.Engine, session *mgo.Session) {
	err := session.DB("edge-local").C("car").Find(nil).All(&car)
	if err != nil {
		panic(err)
	}

	router.GET("/", getData)
	//ここでserverに投稿できるようにrequest設定
	// router.POST()

}

func getData(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"carAll": car,
	})
}

// func postData() {
// 	req, err = http.NewRequest(http.MethodPost, "http://"+url, body)

// 	if err != nil {
// 		fmt.Println("NewRequest error")
// 		return nil, err
// 	}

// }
