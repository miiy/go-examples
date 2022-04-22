package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"sync"
)

type link struct {
	sort int
	title string
	href string
}

func main()  {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var urls []link
	doc.Find("ul.toc a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			log.Println("error link", s)
		}
		urls = append(urls, link{
			sort: i,
			title: s.Text(),
			href: href,
		})
	})

	for k, v := range urls {
		fmt.Println(k, v.sort, v.title, v.href)
	}

	wg := sync.WaitGroup{}
	ch := make(chan int, 8)
	for _, v := range urls {
		ch <- 1
		go func(l link, ch chan int) {
			wg.Add(1)
			defer wg.Done()
			file := "./data/" + strconv.Itoa(l.sort) + ".pdf"

			cmd := exec.Command("wkhtmltopdf",  "https://www.elastic.co/guide/cn/elasticsearch/guide/current/" + l.href, file)
			log.Println(cmd.String())
			err = cmd.Run()
			if err != nil {
				log.Println(cmd.Stderr)
				log.Println(err)
			}

			<- ch
		}(v, ch)
	}

	wg.Wait()
	fmt.Println("Finished.")
}