package main

import (
	"log"
	"net/http"
	"sync"

	tw "testwork/pkg/Data"
	lg "testwork/pkg/RequestsLogic"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go lg.MonitorSites(tw.Sites, &wg)

	http.HandleFunc("/access-time", func(w http.ResponseWriter, r *http.Request) {
		lg.GetUserAccessTime(w, r, tw.Sites)
	})

	http.HandleFunc("/min-access-time", func(w http.ResponseWriter, r *http.Request) {
		lg.GetSiteWithMinAccessTime(w, r, tw.Sites)
	})

	http.HandleFunc("/max-access-time", func(w http.ResponseWriter, r *http.Request) {
		lg.GetSiteWithMaxAccessTime(w, r, tw.Sites)
	})

	http.HandleFunc("/request-counts", func(w http.ResponseWriter, r *http.Request) {
		lg.ShowRequestCounts(w, r)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server started on http://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)

	wg.Wait()
}
