package db

import (
	"fmt"
	"road_violation/db/mongo"
	"time"
)

func NewMongoDB(
	hosts []string, timeout time.Duration, database string, replicaSet string,
	username string, password string, poolLimit int) (mdb *mongo.DB, err error) {

	ds := map[string]interface{}{
		"Use": true,
		"Dn":  "MongoDB",
		"DialInfo": map[string]interface{}{
			"Addrs":          hosts,
			"Timeout":        timeout,
			"Database":       database,
			"ReplicaSetName": replicaSet,
			"Username":       username,
			"Password":       password,
			"PoolLimit":      poolLimit,
		},
	}

	if err = mongo.Setup(ds, true); err != nil {
		return
	}

	mdb, err = mongo.GetDataStore()
	if err != nil {
		return
	}

	err = mdb.Connect()
	if err != nil {
		return
	}

	if !mdb.Connected {
		err = fmt.Errorf("couldn't connect mongodb: %v", mdb.String())
	}

	return
}
