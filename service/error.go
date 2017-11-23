package service

import "net/http"


const (
	StatusNotImplemented = 501
)


func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", StatusNotImplemented)
}


func NotImplementedHandler() http.HandlerFunc {
	return http.HandlerFunc(NotImplemented)
}
