package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	wechat "github.com/OhYee/gowechat"
)

const addr = "0.0.0.0:50005"

type handle struct{}

var wc = wechat.Wechat{
	Token:  "token",
	AESKey: "aes",
}

func (h handle) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	fmt.Printf("URL: %s Params: %+v", req.URL.Path, params)
	switch req.URL.Path {
	case "/CheckSignature":
		if wc.CheckSignature(params.Get("signature"), params.Get("timestamp"), params.Get("nonce")) {
			resp.Write([]byte(params.Get("echostr")))
		} else {
			resp.Write([]byte{})
		}
	default:
		fmt.Printf("Unknow\n")
		resp.WriteHeader(404)
	}

}

func main() {
	err := http.ListenAndServe(addr, handle{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server start at %s\n", addr)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
