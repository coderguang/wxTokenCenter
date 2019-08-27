package wxTokenCenterDef

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgtime"
)

const wx_access_token_key string = "access_token"
const wx_access_token_time_key string = "expires_in"

const wx_access_token_error_code string = "errcode"
const wx_access_token_error_msg string = "errmsg"

type WxConfig struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Appid    string `json:"appid"`
	Secret   string `json:"secret"`
}

type Config struct {
	Port    string     `json:"port"`
	Configs []WxConfig `json:"configs"`
}

type TokenData struct {
	Category         string
	Type             string
	Appid            string
	Secret           string
	RequireFromWx    int //从微信服务器请求的次数
	RequireFromLocal int //本地请求的次数
	TimeoutDt        *sgtime.DateTime
	TokenStr         string
}

type SecureTokenData struct {
	Data map[string](map[string]*TokenData)
	Lock sync.RWMutex
}

func (token *TokenData) ShowMsg() {
	sglog.Info("========start======")
	sglog.Info("Category:%s", token.Category)
	sglog.Info("Type:%s", token.Type)
	sglog.Info("Appid:%s", token.Appid)
	sglog.Info("Secret:%s", token.Secret)
	sglog.Info("RequireFromWx:%d", token.RequireFromWx)
	sglog.Info("RequireFromLocal:%d", token.RequireFromLocal)
	sglog.Info("TimeoutDt:%s", token.TimeoutDt.NormalString())
	sglog.Info("TokenStr:%s", token.TokenStr)
	sglog.Info("========end======")
}

func (token *TokenData) GetAccessTokenFromWx() {

	token.RequireFromWx++

	params := "/cgi-bin/token?grant_type=client_credential&appid=" + token.Appid + "&secret=" + token.Secret

	urls := []string{"api.weixin.qq.com", "api2.weixin.qq.com", "sh.api.weixin.qq.com", "sz.api.weixin.qq.com", "hk.api.weixin.qq.com"}

	success := false
	for _, v := range urls {
		url := "https://" + v + params
		resp, err := http.Get(url)
		if nil != err {
			sglog.Error("get wx access token from %s error,err=%e", url, err)
		} else {
			success = true
			sglog.Info("get wx access token from %s success", url)

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if nil != err {
				sglog.Error("get wx access token error,read resp body error,err=%e", err)
				return
			}
			str := string(body)
			sglog.Info("access_token:%s", str)
			decoder := json.NewDecoder(bytes.NewBufferString(str))
			decoder.UseNumber()
			var result map[string]interface{}
			if err := decoder.Decode(&result); err != nil {
				sglog.Error("json parse failed,str=%s,err=%e", str, err)
				return
			}
			sglog.Info("parse %s json", str)

			if _, ok := result[wx_access_token_error_code]; ok {
				sglog.Error("error token,code=%s", result[wx_access_token_error_code])
				sglog.Error("errmsg=%s", result[wx_access_token_error_msg])
				return
			}

			access_token := result[wx_access_token_key]
			access_token_value, ok := access_token.(string)
			if !ok {
				sglog.Error("parse access_token failed,access_token=%s", access_token)
				return
			}
			sglog.Info("access_token_value:%s", access_token_value)

			time_num := result[wx_access_token_time_key]
			time_num_value, err := time_num.(json.Number).Int64()
			if err != nil {
				sglog.Error("parase time_num failed,time_num=%d", time_num)
				return
			}
			sglog.Info("time:%d", time_num_value)

			token.TokenStr = access_token_value
			token.TimeoutDt = sgtime.New().Add(int(time_num_value))

			sglog.Info("success parse json,token will be timeout at %s", token.TimeoutDt.NormalString())

			break
		}
	}

	if !success {
		sglog.Error("get wx access error from all api,categor:%s,type:%s", token.Category, token.Type)
	}
}
