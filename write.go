package pena

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CircuitLog ...
type CircuitLog struct {
	ID        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Host      string        `json:"host" bson:"host"`
	App       string        `json:"app" bson:"app"`
	Route     string        `json:"route" bson:"route"`
	Fail      bool          `json:"fail" bson:"fail"`
	Tripped   bool          `json:"tripped" bson:"tripped"`
	CreatedAt time.Time     `json:"createdat" bson:"createdat"`
	ErrorCode string        `json:"error_code" bson:"error_code"`
}

// CircuitStatus ...
type CircuitStatus struct {
	Closed  bool
	Fail    bool
	Tripped bool
}

var (
	// DB ...
	DB string
	// Collection ...
	Collection = "service_log"
	// DbURL ..
	DbURL string
)

// SetDB ...
func SetDB(DbPath string, name string) {
	DB = name
	DbURL = DbPath
}

func connectMongo(dbURL string) (*mgo.Session, error) {
	session, err := mgo.Dial(dbURL)
	log.Println(err, "error connect mongo DB")
	session.SetMode(mgo.Monotonic, true)

	return session, err
}

func getCollLog() (*mgo.Session, *mgo.Collection, error) {
	conn, err := connectMongo(DbURL)
	coll := conn.DB(DB).C(Collection)
	return conn, coll, err
}

// WriteLog ...
func WriteLog(circuit CircuitLog) error {
	circuit.CreatedAt = time.Now()
	conn, coll, err := getCollLog()
	err = coll.Insert(circuit)
	log.Println(err)
	conn.Close()
	return err
}
