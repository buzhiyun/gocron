package models

import (
	"bytes"
	"encoding/json"
	"github.com/buzhiyun/gocron/internal/modules/logger"
	"io/ioutil"
	"net/http"
)

type auth struct {
	Id       string `json:"ID"`
	Password string `json:"password"`
}

type loginStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

var PORTAL_API string

func SeptnetAuth(account string, password string) (bool, error) {
	logger.Info("PORTAL_API : ", PORTAL_API)
	//post portal请求提交json数据
	portalUrl := PORTAL_API + "/user/login"
	auths := auth{account, password}
	ba, _ := json.Marshal(auths)
	//logger.Info(string(ba))
	resp, errPortal := http.Post(portalUrl, "application/json", bytes.NewBuffer([]byte(ba)))
	if errPortal != nil {
		logger.Info(errPortal.Error())
		return false, errPortal
	}

	body, _ := ioutil.ReadAll(resp.Body)

	//解析portal返回
	_login := loginStatus{}
	err := json.Unmarshal([]byte(body), &_login)
	if err != nil {
		logger.Info("error:", err)
	}

	logger.Info("登录接口返回 result: ", string(body))
	//fmt.Printf("登录接口返回 result: %s\n", string(body))

	if _login.Status == 200 {
		logger.Info("登录成功")
		return true, nil
	}
	if _login.Status == 403 {
		logger.Info("登录错误")
		return false, nil
	}

	return false, nil
}
