package main

import (
	// "bytes""
	// "io/ioutil"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(request *http.Request) {
		url := *request.URL
		url.Scheme = "https"
		url.Host = "転送したいドメイン"

		// request bodyも転送する場合はこれ
		// reqBody, err := ioutil.ReadAll(request.Body)
		// defer request.Body.Close()
		// if err != nil {
		// 	panic(err)
		// }
		req, err := http.NewRequest(request.Method, url.String(), nil)
		if err != nil {
			panic(err)
		}

		req.Header = request.Header
		*request = *req
	}
	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":18080",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
