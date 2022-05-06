package http

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func GetPort(from, to int) int {
	if to < from {
		panic("to > from")
	}
	return r.Intn(to-from) + from
}

func Serve(addr string, dir string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))

	go func() {
		<-time.After(10 * time.Millisecond)
		log.Printf("Listening %s on http://%s/", dir, addr)
	}()

	log.Fatal(http.ListenAndServe(addr, http.DefaultServeMux))
}
