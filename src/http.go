package main

import (
	"net/http"
)

type MyHttp struct {

}

func HttpRun(ip string, port string) {
	http.HandleFunc("/", http_handler)
	http.ListenAndServe(ip + ":" + port, nil)
}

func http_handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	if rst, err := vm.Run("onRequest('" + r.Form.Encode() + "');"); err != nil {
		log.Error(err.Error());
	} else {
		log.Info(rst)
		w.Write([]byte(rst.String()))

	}

}
