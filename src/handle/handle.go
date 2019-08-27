package wxTokenCenterHandle

import (
	"net/http"
	"strings"
	wxTokenCenterData "wx_common/wx_token_center/src/data"

	"github.com/coderguang/GameEngine_go/sglog"
)

type wx_token_center_handle struct{}

func doGetAccessToken(w http.ResponseWriter, r *http.Request, flag chan bool) {
	r.ParseForm()
	if len(r.Form["key"]) <= 0 {
		w.Write([]byte("")) // not param keys
		sglog.Debug("no key in this handle")
		flag <- true
		return
	}
	rawkeys := r.Form["key"][0]
	keys := strings.Split(rawkeys, ",")
	if len(keys) < 2 {
		w.Write([]byte(""))
		sglog.Debug("require not enough params")
		flag <- true
		return
	}
	tokenstr := wxTokenCenterData.GetAccessToken(keys[0], keys[1])
	w.Write([]byte(tokenstr))
	sglog.Info("get access ok:,categor:%s,type:%s,token:%s", keys[0], keys[1], tokenstr)
	flag <- true
	return
}

func (h *wx_token_center_handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp := make(chan bool)
	go doGetAccessToken(w, r, tmp)
	<-tmp
}

func HttpTokenServer(checkPort string) {
	http.Handle("/", &wx_token_center_handle{})
	port := "0.0.0.0:" + checkPort
	sglog.Info("start require server.listen port:%s", checkPort)
	http.ListenAndServe(port, nil)
}