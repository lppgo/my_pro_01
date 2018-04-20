package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"git.yeeuu.com/yeeuu/mypro_requestURL/models"
)

//访问数据库
func Getdata() {
	datalist, err := models.QueryLocksData()
	if err != nil {
		fmt.Println("查询datalist错误：" + err.Error())
	}

	//=============遍历list并将值请求URL=========
	//初始化一个Http.Client ,初始化客户端是为了设置HTTP请求的参数，例如Cookie
	client := &http.Client{}
	for k, val := range datalist {
		req, err := http.NewRequest("GET", "http://deadpool.yeeuu.com/api/v1/device/info?address="+val.Address, nil)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		//设置Cookie,
		req.Header.Set("Cookie", "u_yeeuu=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IjE4NzY3MTY0NDA3IiwibGV2ZXIiOiJhZG1pbl9jZl9vcF8iLCJyZW1hcmtzIjoiTG91aXNlIiwiZXhwIjoxNTI0OTcxMjU1fQ.hZ42VMJRei993BrB82mOTn148enCIJ_948USyeo7Lm2Edz7RaBr4bNtCesujeNC3MX47a5BOpWIAstbkcBIIJw; account=18767164407")
		res, err := client.Do(req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("body错误：" + err.Error())
			continue
		}
		//解析数据
		var dataInfo models.DeviceBusinessDetailsOut
		err = json.Unmarshal(body, &dataInfo)
		if err != nil {
			fmt.Println("解析数据失败:" + err.Error())
			continue
		}

		// 把请求URL获取的数据写入csv文件  可以利用go协程
		go WriteToFile(dataInfo)
		fmt.Println("-----------------:", k)

	}

}

// 将获取的结构体数据写入csv文件
func WriteToFile(dataInfo models.DeviceBusinessDetailsOut) {
	var lock sync.RWMutex
	lock.Lock()
	defer lock.Unlock()

	f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	// utf8 with bom
	//f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)

	datas := []string{dataInfo.Address, dataInfo.Type, dataInfo.Version, dataInfo.Host, dataInfo.Name, dataInfo.Shop, dataInfo.Partner}

	log.Println(datas)
	w.Write(datas)

	w.Flush()
	f.Close()
}
