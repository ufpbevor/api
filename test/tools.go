package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/ufpblor/api/core"
)

//InitDBTest ...
func InitDBTest() (*mgo.Database, *mgo.Session) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Database: "ufpblor-api",
		Username: "",
		Password: "",
	}
	db := core.GetMongoConnection()
	db.Connect(mongoDBDialInfo)

	return core.GetDatabase()
}

//AssertOk ...
func AssertOk(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

//AssertEquals ...
func AssertEquals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
