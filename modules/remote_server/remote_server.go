package remote_server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/system"

	"github.com/GoAdminGroup/go-admin/modules/logger"
)

const (
	// ServerHost    = "http://localhost:8055"
	// ServerHostApi = "http://localhost:8055/api"

	ServerHost    = "https://www.go-admin.cn"
	ServerHostApi = "https://www.go-admin.cn/api"
)

type LoginRes struct {
	Code int `json:"code"`
	Data struct {
		Token  string `json:"token"`
		Name   string `json:"name"`
		Expire int64  `json:"expire"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func Login(account, password string) LoginRes {
	var resData LoginRes

	req, err := http.NewRequest("POST", ServerHostApi+"/signin", strings.NewReader(`{"account":"`+account+
		`","password":"`+password+`"}`))

	if err != nil {
		logger.Error("login: ", err)
		resData.Code = 500
		resData.Msg = "request error"
		return resData
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		logger.Error("login: ", err)
		resData.Code = 500
		resData.Msg = "request error"
		return resData
	}
	defer func() {
		_ = res.Body.Close()
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("login: ", err)
		resData.Code = 500
		resData.Msg = "request error"
		return resData
	}

	err = json.Unmarshal(body, &resData)
	if err != nil {
		logger.Error("login: ", err)
		resData.Code = 500
		resData.Msg = "request error"
		return resData
	}
	if resData.Code != 0 {
		logger.Error("login to remote GoAdmin server error: ", resData.Msg)
		return resData
	}
	return resData
}

type GetDownloadURLRes struct {
	Code int `json:"code"`
	Data struct {
		Url      string `json:"url"`
		ExtraUrl string `json:"extra_url"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func GetDownloadURL(uuid, token string) (string, string, error) {
	var resData GetDownloadURLRes

	req, err := http.NewRequest("GET", ServerHostApi+"/plugin/download", strings.NewReader(`{"uuid":"`+uuid+`", "version":"`+system.Version()+`"}`))

	if err != nil {
		logger.Error("get plugin download url error: ", err)
		return "", "", err
	}

	req.Header.Add(TokenKey, token)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", "", err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(body, &resData)
	if err != nil {
		return "", "", err
	}
	if resData.Code != 0 {
		return "", "", err
	}
	return resData.Data.Url, resData.Data.ExtraUrl, nil
}

const TokenKey = "GOADMIN_OFFICIAL_SESS"

type GetOnlineReq struct {
	Page       string `json:"page"`
	Free       string `json:"free"`
	PageSize   string `json:"page_size"`
	Filter     string `json:"filter"`
	Order      string `json:"order"`
	Lang       string `json:"lang"`
	CategoryId string `json:"category_id"`
	Version    string `json:"version"`
}

func (req GetOnlineReq) Format() string {
	res := ""
	if req.Page != "" {
		res += "page=" + req.Page + "&"
	}
	if req.PageSize != "" {
		res += "page_size=" + req.PageSize + "&"
	}
	if req.Lang != "" {
		res += "lang=" + req.Lang + "&"
	}
	if req.Filter != "" {
		res += "filter=" + req.Filter + "&"
	}
	if req.Order != "" {
		res += "order=" + req.Order + "&"
	}
	if req.CategoryId != "" {
		res += "category_id=" + req.CategoryId + "&"
	}
	if req.Free != "" {
		res += "free=" + req.Free + "&"
	}
	if req.Version != "" {
		res += "version=" + req.Version + "&"
	}
	if res != "" {
		return res[:len(res)-1]
	}
	return res
}

func GetOnline(reqData GetOnlineReq, token string) ([]byte, error) {
	// TODO: add cache
	req, err := http.NewRequest("GET", ServerHostApi+"/plugin/list?"+reqData.Format(), nil)

	if err != nil {
		logger.Error("get online plugins: ", err)
		return nil, err
	}

	if token != "" {
		req.Header.Add(TokenKey, token)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		logger.Error("get online plugins: ", err)
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("get online plugins: ", err)
		return nil, err
	}

	return body, nil
}
