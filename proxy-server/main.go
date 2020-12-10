package main

import (
	"fmt"
	"github.com/aimamit/hostand/proxy-server/remote"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var (
	InterServerCall = remote.ServerCall{}
)


func main() {
	InterServerCall.Main = remote.ClientInit()

	http.HandleFunc("/", handleRequestAndRedirect)
	log.Fatalln(http.ListenAndServe(":2020", nil))
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	log.Print(req.Host)
	domain := strings.Split(req.Host, ".")
	if len(domain) != 3 {
		log.Fatalln("Not a subdomain based app")
	}

	host, err := InterServerCall.GetIP(domain[0])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
	}
	log.Println(" ", host)
	serveReverseProxy(host, res, req)
}

func serveReverseProxy(ip string, res http.ResponseWriter, req *http.Request) {
	host, _ := url.Parse(fmt.Sprintf("http://%s", ip))

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(host)

	// Update the headers to allow for SSL redirection
	req.URL.Host = host.Host
	req.URL.Scheme = host.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = host.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
