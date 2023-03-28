package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	PREFIX = "/upstream"
)

type HTTPProxy struct {
	proxy *httputil.ReverseProxy
}

func NewHTTPProxy(target string) (*HTTPProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return &HTTPProxy{httputil.NewSingleHostReverseProxy(u)}, nil
}

func (h *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("r:%+v", r)
	h.proxy.ServeHTTP(w, r)
}

func Upstream() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		n, err := io.WriteString(w, fmt.Sprintf("url:%s!\n", r.URL))
		if err != nil {
			log.Fatalf("io.WriteString() failed, err:%s", err)
		}
		log.Printf("io.WriteString() success, n:%d", n)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func main() {
	// 启动Upstream
	// go Upstream()

	proxy, err := NewHTTPProxy("http://127.0.0.1:8088")
	if err != nil {

	}
	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
