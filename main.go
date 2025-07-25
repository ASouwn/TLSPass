package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	target, err := url.Parse("http://localhost:3274")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	// 修改请求 URL 的 Host 和 Scheme，以避免默认使用 https
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = target.Host
	}

	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
	ex, _ := os.Executable()
	certPath := filepath.Join(filepath.Dir(ex), "sccssd.cn.pem")
	keyPath := filepath.Join(filepath.Dir(ex), "sccssd.cn.key")
	err = http.ListenAndServeTLS(":443", certPath, keyPath, nil)
	if err != nil {
		log.Fatal(err)
	}
}
