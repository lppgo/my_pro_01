package models

import (
	"fmt"

	"git.yeeuu.com/yeeuu/mypro_requestURL/db"
	"gopkg.in/mgo.v2/bson"
)

// Device 门锁
type Device struct {
	Host        string `bson:"host"`
	Address     string `bson:"address"`
	Rssi        int    `bson:"rssi"`
	Electric    int    `bson:"electric"`
	Online      bool   `bson:"isOnline"`
	Sn          string `bson:"serial_number"`
	Type        int    `bson:"type"`
	LastSync    int64  `bson:"syncTime"`
	Index       int    `bson:"index"`
	Version     string `bson:"version"`
	LastOnline  int64  `bson:"lastOnline,omitempty"`
	LastOffline int64  `bson:"lastOffline,omitempty"`
	EnterNetAt  int64  `bson:"enter_net"`
	LeaveNetAt  int64  `bson:"leave_net"`
}

// DeviceBusinessDetailsOut 查询业务信息输出
type DeviceBusinessDetailsOut struct {
	Address string `json:"address"`

	Business interface{} `json:"-"`

	PartnerID   string `json:"partnerID,omitempty"`
	Name        string `json:"name,omitempty"`
	Shop        string `json:"shop,omitempty"`
	ApartmentID string `json:"apartmentID,omitempty"`
	ShopID      string `json:"shopID,omitempty"`
	AuthCount   int    `json:"authCount,omitempty"`

	Partner    string `json:"partner,omitempty"`
	Electric   int    `json:"electric"`
	Rssi       int    `json:"rssi"`
	Enternet   int64  `json:"enternet"`
	Sn         string `json:"sn"`
	Type       string `json:"type"` //设备类型
	Version    string `json:"version"`
	Host       string `json:"host"`
	LastOnline int64  `json:"lastTime"`
	LastSync   int64  `json:"lastSync"`
	Online     bool   `json:"online"`
	Updatable  bool   `json:"updatable"`
}

//查询数据库
func QueryLocksData() (data []Device, err error) {
	sess := db.NewDBSession()
	defer sess.Close()

	//mongo 查询语句
	err = sess.DB("yeeuu").C("locks").Find(bson.M{
		"type":    82178,
		"version": bson.M{"$in": []string{"20170615.112311", "20170922.155749"}},
	}).All(&data)
	if err != nil {
		fmt.Println("查询数据库发生错误：" + err.Error())
		fmt.Println(err.Error())
		return nil, err
	}
	return data, err
}
