package mongo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2"
)

type (
	// Mongo ... Global Mongo
	Mongo struct {
		DataStore interface{}
		Setuped   bool
	}

	// DB ... MongoDB session structure
	DB struct {
		Use       bool
		Dn        string
		DialInfo  *mgo.DialInfo
		Session   *mgo.Session
		Connected bool
	}

	// Collection ... Mongodb#Collection
	Collection struct {
		*mgo.Collection
		Session *mgo.Session `json:"-" bson:"-"`
	}
	Option struct {
		Session bool
		DbName  string
		ColName string
	}
)

// DEBUG ... Debug flag
var DEBUG = false

// Global Mgd instance
var mongo = &Mongo{
	DataStore: make(map[string]interface{}),
	Setuped:   false,
}

func Setup(ds map[string]interface{}, autoconnect bool) error {
	if mongo.Setuped {
		return errors.New("Already setup performed")
	}
	if !ds["Use"].(bool) {
		Debug("Skip the datastore. dn=%s\n", ds["Dn"].(string))
		return nil
	}
	mongodb := &DB{}

	err := mapstructure.Decode(ds, mongodb)

	if err != nil {
		return err
	}

	if autoconnect == true {
		Debug("auto-connecting dn=%s\n", ds["Dn"].(string))
		err := mongodb.Connect()
		if err != nil {
			return err
		}
	}
	mongo.DataStore = mongodb
	Debug("Add the datastore. dn=%s\n", mongodb.Dn)

	mongo.Setuped = true
	return nil
}

func (d *DB) Connect() error {
	if d.Connected {
		return nil
	}
	session, err := mgo.DialWithInfo(d.DialInfo)
	if err != nil {
		msg := "Failed mongodb connect."
		Debug("%s", msg)
		return errors.New(msg)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	d.Session = session // Original session

	d.Connected = true

	return err
}

func (d *DB) String() string {

	return fmt.Sprintf("dn=%s, type=MongoDB, connected=%t, addr=%s, database=%s, session=%p",
		d.Dn,
		d.Connected,
		d.DialInfo.Addrs,
		d.DialInfo.Database,
		d.Session,
	)
}

func GetDataStore() (*DB, error) {
	//ver mongoのDatastore
	ds := mongo.DataStore
	if ds == nil {
		return nil, errors.New("Datastore not found")
	}
	if ret, ok := ds.(*DB); ok {
		return ret, nil
	}
	return nil, errors.New("Internal data store type is invalid")
}

func Debug(f string, msgs ...string) {
	if DEBUG {
		fmt.Printf(""+f, strings.Join(([]string)(msgs), " "))
	}
}
