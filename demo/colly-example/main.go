package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type RegionItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {

	// 搜索名称
	searchRegion := "德州市"

	// 目标地区
	var targetRegion *RegionItem

	// 创建 collector
	c := colly.NewCollector(
		// Visit only domains: baike.baidu.com
		colly.AllowedDomains("baike.baidu.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		if targetRegion != nil {
			fmt.Println(e.ChildText("title"), e.Request.URL)
			fmt.Println(e.ChildText(".J-summary"))
		}
	})

	// On every span tag with the class next-button
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if e.Text == "" {
			return
		}

		if e.Text == searchRegion {
			targetRegion = &RegionItem{
				Name: e.Text,
				Url:  e.Name,
			}
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(resp *colly.Response, err error) {
		fmt.Println("OnError", err)
	})

	// Start scraping
	c.Visit("https://baike.baidu.com/item/%E4%B8%AD%E5%8D%8E%E4%BA%BA%E6%B0%91%E5%85%B1%E5%92%8C%E5%9B%BD%E7%9C%81%E7%BA%A7%E8%A1%8C%E6%94%BF%E5%8C%BA/54127472")
	c.Visit("https://baike.baidu.com/item/%E4%B8%AD%E5%8D%8E%E4%BA%BA%E6%B0%91%E5%85%B1%E5%92%8C%E5%9B%BD%E5%9C%B0%E7%BA%A7%E8%A1%8C%E6%94%BF%E5%8C%BA/62589530")
}
