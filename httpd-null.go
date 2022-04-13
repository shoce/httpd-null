/*

GoFmt GoBuild GoRelease
GoRun
ListenAddr=:8080 GoRun

*/

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	LogRequests := os.Getenv("LogRequests")
	ListenAddr := os.Getenv("ListenAddr")
	if ListenAddr == "" {
		ListenAddr = ":80"
	}
	log.Printf("Listening on `%s`.", ListenAddr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if LogRequests != "" {
			log.Printf("%#v\n", r)
		}
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(ListenAddr, nil))
}
