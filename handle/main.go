/*
系统主监听单元(API)，用于开启系统API监听程序，这部分并非必须开启，对于安全性要求较高的环境，可以关闭系统监听
*/
package handle

import (
	"github.com/gorilla/mux"
	"nave/service"
	"net/http"
)

// InfoHandler 获取系统信息
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "token parameter is required", http.StatusBadRequest)
		return
	}

	systemInfo, err := info.GetSystemInfo()
	if err != nil {
		panic(err)
	}

	// Print the JSON formatted system information
	println(string(systemInfo))

	//currentTime := time.Now().Unix()
	response := systemInfo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// 主运行方法
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/info", InfoHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":9998", nil)
}
