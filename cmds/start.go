package cmds

import (
	"bufio"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func Start() {
	proxyMap := proxyInit()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
		}
	})

	for route, proxy := range proxyMap {
		if route == "/" {
			http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != route {
					http.NotFound(w, r)
				}
				proxy.ServeHTTP(w, r)
			})
		} else {
			http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
				proxy.ServeHTTP(w, r)
			})
		}
	}

	http.HandleFunc("/TLSPass", func(w http.ResponseWriter, r *http.Request) {})

	certPath, keyPath := "/etc/TLSPass/tlspass.pem", "/etc/TLSPass/tlspass.key"
	err := http.ListenAndServeTLS(":443", certPath, keyPath, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func readConfig() map[string]string {
	proxyMap, configfile := make(map[string]string), "/etc/TLSPass/config"
	cfg, err := os.Open(configfile)
	if err != nil {
		panic("Failed to open config file:" + err.Error())
	}
	defer cfg.Close()

	scanner := bufio.NewScanner(cfg)
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, ">")
		proxyMap[args[0]] = args[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return proxyMap
}

func proxyInit() map[string]*httputil.ReverseProxy {
	proxyMap := readConfig()
	proxyServerMap := make(map[string]*httputil.ReverseProxy)
	for route, target := range proxyMap {
		targetURL, err := url.Parse(target)
		if err != nil {
			log.Fatalf("Failed to parse target URL %s: %v", target, err)
		}
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		// 修改请求 URL 的 Host 和 Scheme，以避免默认使用 https
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			req.Host = targetURL.Host
		}
		proxyServerMap[route] = proxy
	}
	return proxyServerMap
}
