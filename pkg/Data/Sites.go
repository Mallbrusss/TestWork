package data

import (
	"sync"
	"time"
)

type Site struct {
	Name         string
	URL          string
	LastTime     time.Time
	IsAvailable  bool
	Availability sync.RWMutex
	ResponseTime int64
}

var Sites = []*Site{
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
