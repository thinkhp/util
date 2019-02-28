package thinkHttp

import "net/http"

// 重定向
func ToRedirect(w http.ResponseWriter, newPath string) {
	w.Header().Set("Location", newPath)
	w.WriteHeader(301)
}
