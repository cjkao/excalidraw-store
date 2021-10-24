package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	big "github.com/allegro/bigcache/v3"
)

const (
	ROOM  = "/files/rooms/"
	SHARE = "/files/shareLinks/"
)

type hreqFun func(w http.ResponseWriter, r *http.Request)

func ffHandler(prefix string) hreqFun {
	return func(w http.ResponseWriter, r *http.Request) {
		roomHandler(prefix, w, r)
	}

}
func roomHandler(prefix string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		title := r.URL.Path[len(prefix):]
		body, err := io.ReadAll(r.Body)
		fmt.Printf("save %s --> %d\n", title, len(body))
		if err != nil {
			panic(err)
		}
		cache.Set(title, body)
		w.Write([]byte("ok"))
		return
	}
	if r.Method == "GET" {
		title := r.URL.Path[len(prefix):]

		body, err := cache.Get(title)
		fmt.Printf("load %s --> %d\n", title, len(body))

		if err != nil {
			http.Error(w, "body not found", http.StatusBadRequest)
			return
		}

		w.Write(body)
	}

}

var cache *big.BigCache

// const expire int = 600 // expire in 600 seconds
func main() {
	cache, _ = big.NewBigCache(big.DefaultConfig(10 * time.Minute))

	http.HandleFunc(ROOM, ffHandler(ROOM))
	http.HandleFunc(SHARE, ffHandler(SHARE))
	http.ListenAndServe(":8090", nil)
}
