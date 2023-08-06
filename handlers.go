package main

import "net/http"

func SetupHandler(root string, workers map[string]Worker) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(root)))

	w1 := workers[backends[0]]
	mux.HandleFunc("/1/", func(w http.ResponseWriter, r *http.Request) {
		//w1.BuffReq(r, w)
		w1.handler.ServeHTTP(w, r)
	})

	w2 := workers[backends[1]]
	mux.HandleFunc("/2/", func(w http.ResponseWriter, r *http.Request) {
		//w2.BuffReq(r, w)
		w2.handler.ServeHTTP(w, r)
	})

	w3 := workers[backends[2]]
	mux.HandleFunc("/3/", func(w http.ResponseWriter, r *http.Request) {
		//w3.BuffReq(r, w)
		w3.handler.ServeHTTP(w, r)
	})

	w4 := workers[backends[3]]
	mux.HandleFunc("/4/", func(w http.ResponseWriter, r *http.Request) {
		//w4.BuffReq(r, w)
		w4.handler.ServeHTTP(w, r)
	})

	return mux
}
