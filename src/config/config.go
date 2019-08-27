package wxTokenCenterConfig

import (
	"encoding/json"
	"io/ioutil"
	wxTokenCenterDef "wx_common/wx_token_center/src/define"

	"github.com/coderguang/GameEngine_go/sglog"
)

func ReadConfig(configfile string) *wxTokenCenterDef.Config {
	config, err := ioutil.ReadFile(configfile)
	if err != nil {
		sglog.Fatal("read config error")
		return nil
	}
	t := new(wxTokenCenterDef.Config)
	p := &t
	err = json.Unmarshal([]byte(config), p)
	if err != nil {
		sglog.Fatal("parse config error")
		return nil
	}
	return t
}
