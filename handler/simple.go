package handler

import "net/http"

type Simple func(w http.ResponseWriter, r *http.Request)

func (h Simple) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func NewSimpleHandler(responseCode int, contentType string, responseBody []byte) Simple {
	return Simple(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contentType)
		w.WriteHeader(responseCode)
		w.Write(responseBody)
	})
}
