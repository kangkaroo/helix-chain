package api

import (
	"fmt"
	"net/http"
)

// 处理 API 请求
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Consortium Blockchain API")
}
