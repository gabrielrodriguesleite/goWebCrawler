package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gabrielrodriguesleite/goWebCrawler/db"
	"github.com/gabrielrodriguesleite/goWebCrawler/website"
	"golang.org/x/net/html"
)

var (
	link   string
	action string
)

func init() {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br", "url para iniciar visitas")
	flag.StringVar(&action, "action", "website", "qual serviço usar")
}

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

func main() {
	flag.Parse()
	fmt.Println("Web Crawler Go v1.0.0")

	switch action {
	case "website":
		website.Run()
	case "webcrawler":
		done := make(chan bool)
		go visitLink(link)
		<-done
	default:
		fmt.Printf("Action não reconhecida: %s\n", action)
	}
}

func visitLink(url string) {
	fmt.Printf("Visitando: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Não foi possível acessar: %s\nErro: %v", url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[Erro] Status code diferente de 200: %v\n", resp.StatusCode)
	}

	element, err := html.Parse(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Erro ao traduzir html: %v", err))
	}

	extractLinks(element)
}

func extractLinks(element *html.Node) {
	if element.Type == html.ElementNode && element.Data == "a" {
		for _, attr := range element.Attr {
			if attr.Key != "href" {
				continue // continua buscando até achar href
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme != "https" {
				continue
			}

			if db.VisitedLink(link.String()) {
				fmt.Printf("link já visitado: %s\n", link)
				continue
			}

			visitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}

			db.Insert("links", visitedLink)

			go visitLink(link.String())
		}
	}

	for c := element.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}
