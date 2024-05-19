package server

import "net/http"

func CreateServer(address string) *http.Server {
	s := &http.Server{
		Addr: address,
	}
	return s
}
