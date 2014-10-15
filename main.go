package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"github.com/quirkey/magick"
)

var username = ""
var password = ""
var listenOn = "127.0.0.1:7070"

type imageFuzzer struct{}

func (i *imageFuzzer) Handle(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if resp == nil {
		return resp
	}

	var ext string
	l, err := resp.Location()
	if err != nil {
		if ctx == nil || ctx.Req == nil {
			return resp
		}
		parts := strings.Split(ctx.Req.URL.Path, ".")
		ext = strings.ToLower(parts[len(parts)-1])
	} else {
		parts := strings.Split(l.Path, ".")
		ext = strings.ToLower(parts[len(parts)-1])
	}

	switch strings.ToLower(resp.Header.Get("Content-Type")) {
	case "image/gif":
		ext = "gif"
	case "image/jpg", "image/jpeg":
		ext = "jpg"
	}

	switch ext {
	case "gif", "jpg", "jpeg", "png":
		raw := new(bytes.Buffer)
		raw.ReadFrom(resp.Body)
		resp.Body.Close()
		data, err := magick.NewFromBlob(raw.Bytes(), ext)

		if err != nil {
			resp.Body = ioutil.NopCloser(raw)
			return resp
		}

		size := data.Width()
		if size < data.Height() {
			size = data.Height()
		}

		oSize := size
		if oSize > 1024 {
			data.Resize("1024")
		}
		for i := 0; i < 4; i++ {
			data.Resize("20%")
			data.Resize("500%")
		}
		if oSize > 1024 {
			if data.Width() == oSize {
				data.Resize(fmt.Sprintf("%d", oSize))
			} else {
				data.Resize(fmt.Sprintf("x%d", oSize))
			}
		}

		newBytes, err := data.ToBlob(ext)
		if err != nil {
			resp.Body = ioutil.NopCloser(raw)
			return resp
		}
		raw.Truncate(0)
		raw.Write(newBytes)
		resp.Body = ioutil.NopCloser(raw)
		return resp
	default:
		return resp
	}
}

var fuzzer = &imageFuzzer{}

func init() {
	flag.StringVar(&username, "u", username, "Proxy username.  Empty string disables authentication.")
	flag.StringVar(&password, "p", password, "Proxy password")
	flag.StringVar(&listenOn, "l", listenOn, "Proxy listening address")
}

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	if username != "" {
		auth.ProxyBasic(proxy, "fuzzyproxy", func(user, pass string) bool {
			if user == username && pass == password {
				return true
			}
			return false
		})
	}
	proxy.OnResponse().Do(fuzzer)
	log.Fatal(http.ListenAndServe(listenOn, proxy))
}
