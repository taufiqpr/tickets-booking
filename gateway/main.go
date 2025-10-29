package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"ticket-booking/gateway/config"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("load env: %v", err)
	}
	mux := http.NewServeMux()

	mux.Handle("/users/", makeProxy(cfg.UserBaseURL, "/users/"))
	mux.Handle("/trains/", makeProxy(cfg.TrainBaseURL, "/trains/"))
	mux.Handle("/schedules/", makeProxy(cfg.SchedBaseURL, "/schedules/"))
	mux.Handle("/bookings/", makeProxy(cfg.BookBaseURL, "/bookings/"))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })

	log.Printf("gateway listening on %s", cfg.Addr())
	log.Fatal(http.ListenAndServe(cfg.Addr(), mux))
}

func makeProxy(targetBase, prefix string) http.Handler {
	t, _ := url.Parse(targetBase)
	rp := httputil.NewSingleHostReverseProxy(t)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		if !strings.HasPrefix(r.URL.Path, "/") { r.URL.Path = "/" + r.URL.Path }
		rp.ServeHTTP(w, r)
	})
}
