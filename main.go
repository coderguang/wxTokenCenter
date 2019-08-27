package main

import (
	"log"
	"os"
	wxTokenCenterData "wxTokenCenter/src/data"
	wxTokenCenterHandle "wxTokenCenter/src/handle"

	"github.com/coderguang/GameEngine_go/sgcmd"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func ShowAllTokenMsg(cmd []string) {
	wxTokenCenterData.ShowAllTokeData()
}

func RegistCmd() {
	// ["ShowAllTokenMsg"]
	sgcmd.RegistCmd("ShowAllTokenMsg", "[\"ShowAllTokenMsg\"]", ShowAllTokenMsg)
}

func main() {

	logPath := "./log/"
	sgserver.StartLogServer("debug", logPath, log.LstdFlags, true)

	arg_num := len(os.Args) - 1

	if arg_num < 1 {
		sglog.Fatal("please input config file")
		return
	}
	configfile := os.Args[1]

	wxTokenCenterData.InitConfig(configfile)

	wxTokenCenterData.InitTokenData()

	go wxTokenCenterHandle.HttpTokenServer(wxTokenCenterData.GetListenPort())

	RegistCmd()
	sgcmd.StartCmdWaitInputLoop()

}
