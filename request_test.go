package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func indexRequest(w http.ResponseWriter, r *http.Request) {
}

func BenchmarkReqBuffer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytes.Buffer{}
			enc := json.NewEncoder(&b)
			if err := enc.Encode(obj); err != nil {
				panic(err)
			}
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/request", &b); err != nil {
				panic(err)
			} else if res, err := client.Do(req); err != nil {
				panic(err)
			} else {
				_, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				res.Body.Close()
			}
		}
	})
}

func BenchmarkReqMarshall(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf, err := json.Marshal(obj)
			if err != nil {
				panic(err)
			}
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/request", bytes.NewReader(buf)); err != nil {
				panic(err)
			} else if res, err := client.Do(req); err != nil {
				panic(err)
			} else {
				_, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				res.Body.Close()
			}
		}
	})
}

func BenchmarkReqPipe(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r, w := io.Pipe()
			enc := json.NewEncoder(w)

			go func() {
				if err := enc.Encode(obj); err != nil {
					panic(err)
				}
				w.Close()
			}()

			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/request", r); err != nil {
				panic(err)
			} else if res, err := client.Do(req); err != nil {
				panic(err)
			} else {
				_, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				res.Body.Close()
			}
		}
	})
}
