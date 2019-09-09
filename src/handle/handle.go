package wxTokenCenterHandle

import (
	"net/http"
	"strings"
	wxTokenCenterData "wxTokenCenter/src/data"

	"github.com/coderguang/GameEngine_go/sglog"
)

type wx_token_center_handle struct{}

func doGetAccessToken(w http.ResponseWriter, r *http.Request, flag chan bool) {

	defer func() {
		flag <- true
	}()

	r.ParseForm()
	if len(r.Form["key"]) <= 0 {
		w.Write([]byte("{\"errcode\":3}")) // not param keys
		sglog.Debug("no key in this handle")

		return
	}
	rawkeys := r.Form["key"][0]
	keys := strings.Split(rawkeys, ",")
	if len(keys) < 2 {
		w.Write([]byte("{\"errcode\":4}"))
		sglog.Debug("require not enough params")

		return
	}
	tokenstr := wxTokenCenterData.GetAccessToken(keys[0], keys[1])
	w.Write([]byte(tokenstr))
	sglog.Info("get access ok:,categor:%s,type:%s,token:%s", keys[0], keys[1], tokenstr)

	return
}

func (h *wx_token_center_handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp := make(chan bool)
	go doGetAccessToken(w, r, tmp)
	<-tmp
	close(tmp)
}

func HttpTokenServer(checkPort string) {
	http.Handle("/", &wx_token_center_handle{})
	port := "0.0.0.0:" + checkPort
	sglog.Info("start require server.listen port:%s", checkPort)
	http.ListenAndServe(port, nil)
}
