package main

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/erikdubbelboer/gojsonspeed/openrtb"
	"github.com/julienschmidt/httprouter"
)

const JSON = `
    {
        "at": 2,
        "cur": [
            "USD"
        ],
        "device": {
            "carrier": "Unknown",
            "connectiontype": 3,
            "devicetype": 7,
            "geo": {
                "country": "USA",
                "type": 2
            },
            "ip": "107.167.108.35",
            "js": 1,
            "os": "Android",
            "ua": "Opera/9.80 (Android; Opera Mini/9.0.1829/37.6334; U; en) Presto/2.12.423 Version/12.16"
        },
        "id": "VdwCt9kBQf9cVVG_ltcpAw",
        "imp": [
            {
                "banner": {
                    "h": 0,
                    "id": "1",
                    "w": 0
                },
                "bidfloorcur": "USD",
                "id": "1",
                "tagid": "7928"
            }
        ],
        "site": {
            "cat": [
                "IAB10-2"
            ],
            "domain": "hqhotwallpaper.blogspot.in",
            "id": "8033",
            "name": "http://hqhotwallpaper.blogspot.in/",
            "publisher": {
                "id": "7294",
                "name": "Monikandan R"
            },
            "ref": "http://hqhotwallpaper.blogspot.in/2012/03/cinemaki-veladam-randi-movie-latest-hot.html?m=1"
        },
        "test": 0,
        "tmax": 100,
        "user": {
            "id": "VdwCq1MB8EcSnFoQR6pKog"
        }
    }
`

var (
	obj openrtb.BidRequest

	client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			DisableCompression:  false,
			MaxIdleConnsPerHost: 500,
		},
		Jar:     nil,
		Timeout: time.Millisecond * 100,
	}
)

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func notfound(w http.ResponseWriter, r *http.Request) {
	panic("404")
}

func init() {
	router := httprouter.New()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false
	router.NotFound = http.HandlerFunc(notfound)

	router.HandlerFunc("POST", "/readunmarshal", indexReadUnmarshal)
	router.HandlerFunc("POST", "/readdecoder", indexReadDecoder)
	router.HandlerFunc("POST", "/respmarshal", indexRespMarshal)
	router.HandlerFunc("POST", "/respencoder", indexRespEncoder)
	router.HandlerFunc("POST", "/respbuffer", indexRespBuffer)
	router.HandlerFunc("POST", "/request", indexRequest)

	s := &http.Server{
		Addr:           "127.0.0.1:12346",
		Handler:        router,
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 1024 * 6, // 4k cookie + 2k http request
	}

	s.SetKeepAlivesEnabled(true)

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := s.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
			panic(err)
		}
	}()

	if err := json.Unmarshal([]byte(JSON), &obj); err != nil {
		panic(err)
	}
}
