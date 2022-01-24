package handler

import "net/http"

type Simple func(w http.ResponseWriter, r *http.Request)

func (h Simple) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func NewSimpleHandler(responseCode int, responseBody []byte) Simple {
	return Simple(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		w.Write(responseBody)
	})
}
