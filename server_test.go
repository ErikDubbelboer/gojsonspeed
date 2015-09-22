package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/erikdubbelboer/gojsonspeed/openrtb"
)

func indexReadDecoder(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)

	m := openrtb.BidRequest{}

	if err := dec.Decode(&m); err != nil {
		panic(err)
	}
}

func indexRespEncoder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err := enc.Encode(obj); err != nil {
		panic(err)
	}
}

func indexReadUnmarshal(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	m := openrtb.BidRequest{}

	if err := json.Unmarshal(body, &m); err != nil {
		panic(err)
	}
}

func indexRespBuffer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	b := bytes.Buffer{}
	enc := json.NewEncoder(&b)
	if err := enc.Encode(obj); err != nil {
		panic(err)
	}
	b.WriteTo(w)
}

func indexRespMarshal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if buf, err := json.Marshal(obj); err != nil {
		panic(err)
	} else {
		w.Write(buf)
	}
}

func BenchmarkReadUnmarshal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/readunmarshal", bytes.NewReader([]byte(JSON))); err != nil {
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

func BenchmarkReadDecoder(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/readdecoder", bytes.NewReader([]byte(JSON))); err != nil {
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

func BenchmarkRespEncoder(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/respencoder", bytes.NewReader([]byte(JSON))); err != nil {
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

func BenchmarkRespBuffer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/respbuffer", bytes.NewReader([]byte(JSON))); err != nil {
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

func BenchmarkRespMarshal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if req, err := http.NewRequest("POST", "http://127.0.0.1:12346/respmarshal", bytes.NewReader([]byte(JSON))); err != nil {
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
