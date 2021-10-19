package on

import (
	"io"
	"os"
	"road_violation/db"
	"road_violation/db/mongo"
	"time"
)

const LocalEnv = "local"

type Local struct {
	*mongo.DB
	accessLogger   io.Writer
	activityLogger io.Writer
	goEnv          string
	port           int
	wsPort         int
}

// Init is a initialize method　空のlocalが渡されてる
func (lo *Local) Init() (err error) {
	// GoEnv and port
	lo.goEnv = LocalEnv
	lo.port = 3000
	lo.wsPort = 3001

	// Loggers
	lo.accessLogger = os.Stdout
	lo.activityLogger = os.Stdout

	// MongoDB
	mdb, err := db.NewMongoDB(
		//ここlocalhostにすればいけるのでは?
		[]string{"db:27017"},
		10*time.Second,
		"edge-local",
		"",
		"",
		"",
		128,
	)

	if err != nil {
		return err
	}

	lo.DB = mdb
	return
}

func (lo *Local) GetDB() *mongo.DB {
	return lo.DB
}
