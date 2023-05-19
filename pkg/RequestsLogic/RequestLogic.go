package requestslogic

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	tw "testwork/pkg/Data"
)

var requestCountMap = map[string]int{
	"/access-time":     0,
	"/min-access-time": 0,
	"/max-access-time": 0,
}

func RequestCount(endpoint string) {
	if count, ok := requestCountMap[endpoint]; ok {
		requestCountMap[endpoint] = count + 1
	}
}

func MonitorSites(sites []*tw.Site, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		for _, site := range sites {
			go func(site *tw.Site) {
				resp, err := http.Head(site.URL)
				available := false
				if err == nil && resp.StatusCode == http.StatusOK {
					available = true
				}

				site.Availability.Lock()
				site.IsAvailable = available
				site.LastTime = time.Now()
				site.Availability.Unlock()
			}(site)
		}

		time.Sleep(time.Minute)
	}
}

func GetUserAccessTime(w http.ResponseWriter, r *http.Request, sites []*tw.Site) {
	RequestCount("/access-time")
	siteName := r.URL.Query().Get("site")

	for _, site := range sites {
		if site.Name == siteName {
			site.Availability.RLock()
			lastTime := site.LastTime
			site.Availability.RUnlock()

			fmt.Fprint(w, lastTime.String())
			return
		}
	}

	http.Error(w, "Site not found", http.StatusNotFound)
}

func GetSiteWithMinAccessTime(w http.ResponseWriter, r *http.Request, sites []*tw.Site) {
	RequestCount("/min-access-time")
	if len(sites) == 0 {
		http.Error(w, "No sites available", http.StatusNotFound)
		return
	}

	minSite := sites[0]
	minSite.Availability.RLock()
	minLastTime := minSite.LastTime
	minSite.Availability.RUnlock()

	for _, site := range sites {
		site.Availability.RLock()
		lastTime := site.LastTime
		site.Availability.RUnlock()

		if lastTime.Before(minLastTime) {
			minSite = site
			minLastTime = lastTime
		}
	}
	fmt.Fprint(w, minSite.Name)
	
}

func GetSiteWithMaxAccessTime(w http.ResponseWriter, r *http.Request, sites []*tw.Site) {
	RequestCount("/max-access-time")
	if len(sites) == 0 {
		http.Error(w, "No sites available", http.StatusNotFound)
		return
	}

	maxSite := sites[0]
	maxSite.Availability.RLock()
	maxLastTime := maxSite.LastTime
	maxSite.Availability.RUnlock()

	for _, site := range sites {
		site.Availability.RLock()
		lastTime := site.LastTime
		site.Availability.RUnlock()

		if lastTime.After(maxLastTime) {
			maxSite = site
			maxLastTime = lastTime
		}
	}

	fmt.Fprintf(w, maxSite.Name)
}

func ShowRequestCounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	for endpoint, count := range requestCountMap {
		fmt.Fprintf(w, "<p>%s: %d</p>", endpoint, count)
	}
}
