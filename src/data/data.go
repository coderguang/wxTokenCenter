package wxTokenCenterData

import (
	wxTokenCenterConfig "wxTokenCenter/src/config"
	wxTokenCenterDef "wxTokenCenter/src/define"

	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sgtime"

	"github.com/coderguang/GameEngine_go/sglog"
)

var globalcfg *wxTokenCenterDef.Config
var globalTokenDataMap *wxTokenCenterDef.SecureTokenData

func InitConfig(configfile string) {
	globalcfg = wxTokenCenterConfig.ReadConfig(configfile)
}

func GetListenPort() string {
	return globalcfg.Port
}

func InitTokenData() {
	globalTokenDataMap = new(wxTokenCenterDef.SecureTokenData)
	globalTokenDataMap.Data = make(map[string](map[string]*wxTokenCenterDef.TokenData))

	globalTokenDataMap.Lock.Lock()
	defer globalTokenDataMap.Lock.Unlock()
	for _, v := range globalcfg.Configs {

		if typeMap, ok := globalTokenDataMap.Data[v.Category]; ok {
			if _, okex := typeMap[v.Type]; okex {
				sglog.Error("duplate token config,category:%s,type:%s", v.Category, v.Type)
				sgthread.DelayExit(2)
			} else {
				tmp := new(wxTokenCenterDef.TokenData)
				tmp.Category = v.Category
				tmp.Type = v.Type
				tmp.Appid = v.Appid
				tmp.Secret = v.Secret
				tmp.RequireFromLocal = 0
				tmp.RequireFromWx = 0
				tmp.TimeoutDt = sgtime.New()
				tmp.TokenStr = "{\"errcode\":1}"
				tmp.GetAccessTokenFromWx()
				typeMap[v.Type] = tmp
			}
		} else {

			tmp := new(wxTokenCenterDef.TokenData)
			tmp.Category = v.Category
			tmp.Type = v.Type
			tmp.Appid = v.Appid
			tmp.Secret = v.Secret
			tmp.RequireFromLocal = 0
			tmp.RequireFromWx = 0
			tmp.TimeoutDt = sgtime.New()
			tmp.TokenStr = "{\"errcode\":1}"
			tmp.GetAccessTokenFromWx()
			typeMap = make(map[string]*wxTokenCenterDef.TokenData)
			typeMap[v.Type] = tmp
			globalTokenDataMap.Data[v.Category] = typeMap
		}

	}

	sglog.Info("init token data success:size=%d", len(globalcfg.Configs))
}

func GetAccessToken(Category string, Type string) string {

	globalTokenDataMap.Lock.Lock()
	defer globalTokenDataMap.Lock.Unlock()
	if typeMap, ok := globalTokenDataMap.Data[Category]; ok {
		if v, okex := typeMap[Type]; okex {
			now := sgtime.New()
			v.RequireFromLocal++
			if now.Before(v.TimeoutDt) {
				return v.TokenStr
			} else {
				v.GetAccessTokenFromWx()
				return v.TokenStr
			}
		}
	}
	sglog.Error("no this config,category:%s,type:%s", Category, Type)
	return "{\"errcode\":2}"
}

func ShowAllTokeData() {
	sglog.Info("------------start show all-----------")
	globalTokenDataMap.Lock.Lock()
	defer globalTokenDataMap.Lock.Unlock()
	for _, v := range globalTokenDataMap.Data {
		for _, vv := range v {
			vv.ShowMsg()
		}
	}

	sglog.Info("------------end show all-----------")
}
