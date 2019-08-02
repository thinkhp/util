package thinkHttp

import (
	"net"
	"net/http"
)
const localhost = "127.0.0.1"
// nginx
//proxy_set_header Host $host;
//proxy_set_header X-Real-IP $remote_addr;
//proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
func GetRealIp(r *http.Request) (realIp string){
	addr := r.RemoteAddr
	//fmt.Println(r.Proto, r.Host, r.Form, realIp, addr)
	host, _, _ := net.SplitHostPort(addr)
	realIp = host
	if host == "127.0.0.1" || host == "::1" || host == "localhost" {
		// 检查是否因为代理
		if r.Header.Get("X-Real-IP") != "" {
			realIp = r.Header.Get("X-Real-IP")
		} else {
			realIp = localhost
		}
	}

	return realIp
}
