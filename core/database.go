package core

import (
	"crypto/tls"
	"fmt"
	"net"

	"gopkg.in/mgo.v2"
)

var conn *MongoConnection
var mongoDBDialInfo *mgo.DialInfo

func init() {
	conn = new(MongoConnection)
}

//GetDialInfo ...
func GetDialInfo(host string, ssl bool) *mgo.DialInfo {
	dialInfo, err := mgo.ParseURL(host)
	if err != nil {
		fmt.Println("Erro: Parse URL host mongo.")
		panic(err)
	}

	if ssl {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		}
	}

	mongoDBDialInfo = dialInfo
	return mongoDBDialInfo
}

//GetMongoConnection Returns the current global connection object
func GetMongoConnection() *MongoConnection {
	return conn
}

//GetDatabase Returns a the project database and the current session
func GetDatabase() (*mgo.Database, *mgo.Session) {
	s := conn.session.Copy()
	return s.DB(mongoDBDialInfo.Database), s
}

//MongoConnection Stores the global mongo session
type MongoConnection struct {
	session *mgo.Session
}

// Connect to the dabase with given user and password and store the session
func (m *MongoConnection) Connect(DialInfo *mgo.DialInfo) {
	session, err := mgo.DialWithInfo(DialInfo)
	if err != nil {
		fmt.Println("Erro: Authentication failed!")
		panic(err)
	}

	mongoDBDialInfo = DialInfo

	session.SetMode(mgo.Monotonic, true)
	m.session = session
	fmt.Printf("Conected.\n")
}

// Close the stored session
func (m *MongoConnection) Close() {
	m.session.Close()
}
