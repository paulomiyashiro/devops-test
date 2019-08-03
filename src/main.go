package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)
func request(w http.ResponseWriter, r *http.Request) {
	points := rand.Intn(40)
	var m string
	if points > 1 {
		m = "points"
	} else {
		m = "point"
	}
	fmt.Fprintf(w, "Hey! You win "+strconv.Itoa(points)+" "+m)
}

func main() {

	http.HandleFunc("/", request)
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	http.ListenAndServeTLS("0.0.0.0:8443", "/etc/ssl/certs/server.crt", "/etc/ssl/private/server.key", nil)
}
