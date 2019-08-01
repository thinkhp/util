package thinkHttp

import (
	"net/http"
	"os"
	"testing"
	"util/think"
)

func Test1(t *testing.T){
	port := ":8081"
	//dir, err := os.Getwd()
	//think.IsNil(err)
	http.HandleFunc("/", PageHandler)

	err := http.ListenAndServe(port, nil)
	think.IsNil(err)
}

func PageHandler(w http.ResponseWriter, r *http.Request){
	WriteHTMLPage(w, "./static/test.html")
}


func Test2(t *testing.T){
	port := ":8080"
	dir, err := os.Getwd()
	think.IsNil(err)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static"))))

	err = http.ListenAndServe(port, nil)
	think.IsNil(err)
}
