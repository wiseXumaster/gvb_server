package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/common/res"
	"gvb_server/core"
	"gvb_server/global"
	"os"
)

const file = "common/res/err_code.json"

//type ErrResponse struct {
//}

//type ErrMap map[string]string
type ErrMap map[res.ErrorCode]string

func main() {
	// 读取配置文件
	core.InitConf()
	global.Log = core.InitLogger()
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	//fmt.Println(errMap)
	fmt.Println(errMap[res.SettingsError])
}
