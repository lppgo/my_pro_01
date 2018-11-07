package db

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var (
	sess *mgo.Session
)

//=================连接mongoDB数据库，并设置数据库参数===================
func InitDB() error {
	var err error
	sess, err = mgo.Dial("mongodb://username:pw@ip:port")
	if err != nil {
		fmt.Println("--------连接数据库异常-：-------")
		fmt.Println(err.Error())
		return err
	}

	if err := sess.Ping(); err != nil {
		fmt.Println("数据库ping失败！")
		return err
	}

	sess.SetSocketTimeout(5 * time.Second)
	sess.SetSyncTimeout(5 * time.Second)
	sess.SetPoolLimit(300)
	sess.SetMode(mgo.Monotonic, true)
	return nil
}
func CheckStatus() bool {
	return sess.Ping() == nil
}
func NewDBSession() *mgo.Session {
	return sess.Copy()
}
