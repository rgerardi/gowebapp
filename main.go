package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s: %s - %s - %s", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"), r.Method, r.UserAgent(), r.URL)
	log.Info(message)
	fmt.Fprintln(w, "There's an API here")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s: %s - %s - %s", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"), r.Method, r.UserAgent(), r.URL)
	log.Info(message)
	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
