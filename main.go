package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Site struct {
	Name         string
	URL          string
	LastTime     time.Time
	IsAvailable  bool
	Availability sync.RWMutex
}

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

func MonitorSites(sites []*Site, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		for _, site := range sites {
			go func(site *Site) {
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

func GetUserAccessTime(w http.ResponseWriter, r *http.Request, sites []*Site) {
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

func GetSiteWithMinAccessTime(w http.ResponseWriter, r *http.Request, sites []*Site) {
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

	fmt.Fprintf(w, minSite.Name)
}

func GetSiteWithMaxAccessTime(w http.ResponseWriter, r *http.Request, sites []*Site) {
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
	fmt.Fprintf(w, "<h1>Request Counts</h1>")
	for endpoint, count := range requestCountMap {
		fmt.Fprintf(w, "<p>%s: %d</p>", endpoint, count)
	}
}

func ShowAdminPage(w http.ResponseWriter, r *http.Request) {
    // Serve the admin.html file
    http.ServeFile(w, r, "admins.html")
}

func main() {
	sites := []*Site{
		{Name: "Google", URL: "https://www.google.com"},
		{Name: "YouTube", URL: "https://www.youtube.com"},
		{Name: "Facebook", URL: "https://www.facebook.com"},
		{Name: "Baidu", URL: "https://www.baidu.com"},
		{Name: "Wikipedia", URL: "https://www.wikipedia.org"},
		{Name: "QQ", URL: "https://www.qq.com"},
		{Name: "Taobao", URL: "https://www.taobao.com"},
		{Name: "Yahoo", URL: "https://www.yahoo.com"},
		{Name: "Tmall", URL: "https://www.tmall.com"},
		{Name: "Amazon", URL: "https://www.amazon.com"},
		{Name: "Google India", URL: "https://www.google.co.in"},
		{Name: "Twitter", URL: "https://www.twitter.com"},
		{Name: "Sohu", URL: "https://www.sohu.com"},
		{Name: "JD", URL: "https://www.jd.com"},
		{Name: "Microsoft Live", URL: "https://www.live.com"},
		{Name: "Instagram", URL: "https://www.instagram.com"},
		{Name: "Sina", URL: "https://www.sina.com.cn"},
		{Name: "Weibo", URL: "https://www.weibo.com"},
		{Name: "Google Japan", URL: "https://www.google.co.jp"},
		{Name: "Reddit", URL: "https://www.reddit.com"},
		{Name: "VK", URL: "https://www.vk.com"},
		{Name: "360", URL: "https://www.360.cn"},
		{Name: "Tmall Login", URL: "https://login.tmall.com"},
		{Name: "Blogspot", URL: "https://www.blogspot.com"},
		{Name: "Yandex", URL: "https://www.yandex.ru"},
		{Name: "Google Hong Kong", URL: "https://www.google.com.hk"},
		{Name: "Netflix", URL: "https://www.netflix.com"},
		{Name: "LinkedIn", URL: "https://www.linkedin.com"},
		{Name: "Pornhub", URL: "https://www.pornhub.com"},
		{Name: "Google Brazil", URL: "https://www.google.com.br"},
		{Name: "Twitch", URL: "https://www.twitch.tv"},
		{Name: "Tmall Pages", URL: "https://pages.tmall.com"},
		{Name: "CSDN", URL: "https://www.csdn.net"},
		{Name: "Yahoo Japan", URL: "https://www.yahoo.co.jp"},
		{Name: "Mail.ru", URL: "https://www.mail.ru"},
		{Name: "AliExpress", URL: "https://www.aliexpress.com"},
		{Name: "Alipay", URL: "https://www.alipay.com"},
		{Name: "Office", URL: "https://www.office.com"},
		{Name: "Google France", URL: "https://www.google.fr"},
		{Name: "Google Russia", URL: "https://www.google.ru"},
		{Name: "Google UK", URL: "https://www.google.co.uk"},
		{Name: "Microsoft Online", URL: "https://www.microsoftonline.com"},
		{Name: "Google Germany", URL: "https://www.google.de"},
		{Name: "eBay", URL: "https://www.ebay.com"},
		{Name: "Microsoft", URL: "https://www.microsoft.com"},
		{Name: "LiveJasmin", URL: "https://www.livejasmin.com"},
		{Name: "Twitter Shortener", URL: "https://t.co"},
		{Name: "Bing", URL: "https://www.bing.com"},
		{Name: "Xvideos", URL: "https://www.xvideos.com"},
		{Name: "Google Canada", URL: "https://www.google.ca"},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go MonitorSites(sites, &wg)

	http.HandleFunc("/access-time", func(w http.ResponseWriter, r *http.Request) {
		GetUserAccessTime(w, r, sites)
	})

	http.HandleFunc("/min-access-time", func(w http.ResponseWriter, r *http.Request) {
		GetSiteWithMinAccessTime(w, r, sites)
	})

	http.HandleFunc("/max-access-time", func(w http.ResponseWriter, r *http.Request) {
		GetSiteWithMaxAccessTime(w, r, sites)
	})

	http.HandleFunc("/request-counts", ShowRequestCounts)
    
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("admins", ShowAdminPage)
	

	log.Println("Server started on http://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)

	wg.Wait()
}
