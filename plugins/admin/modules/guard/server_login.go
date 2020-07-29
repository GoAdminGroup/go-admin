package guard

import (
	"encoding/json"
	"io/ioutil"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/logger"
)

type ServerLoginParam struct {
	Account  string
	Password string
}

func (g *Guard) ServerLogin(ctx *context.Context) {

	var p ServerLoginParam

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		logger.Error("get server login parameter error: ", err)
	}

	err = json.Unmarshal(body, &p)

	if err != nil {
		logger.Error("unmarshal server login parameter error: ", err)
	}

	ctx.SetUserValue(serverLoginParamKey, &p)
	ctx.Next()
}

func GetServerLoginParam(ctx *context.Context) *ServerLoginParam {
	return ctx.UserValue[serverLoginParamKey].(*ServerLoginParam)
}
