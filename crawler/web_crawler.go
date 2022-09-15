package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gabrielrodriguesleite/goWebCrawler/db"
	"golang.org/x/net/html"
)

type webCrawler struct {
	log chan string
}

func (wc *webCrawler) VisitLink(url string) {
	wc.log <- fmt.Sprintf("Visitando: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Não foi possível acessar: %s\nErro: %v", url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		wc.log <- fmt.Sprintf("[Erro] Status code diferente de 200: %v", resp.StatusCode)
	}

	element, err := html.Parse(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Erro ao traduzir html: %v", err))
	}

	wc.extractLinks(element)
}

func (wc *webCrawler) extractLinks(element *html.Node) {
	if element.Type == html.ElementNode && element.Data == "a" {
		for _, attr := range element.Attr {
			if attr.Key != "href" {
				continue // continua buscando até achar href
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme != "https" {
				continue
			}

			if db.CheckVisitedLink(link.String()) {
				wc.log <- fmt.Sprintf("link já visitado: %s", link)
				continue
			}

			visitedLink := db.VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}

			db.Insert("links", visitedLink)

			go wc.VisitLink(link.String())
		}
	}

	for c := element.FirstChild; c != nil; c = c.NextSibling {
		wc.extractLinks(c)
	}
}

func (wc *webCrawler) Log() chan string { return wc.log }

func New() *webCrawler { return &webCrawler{log: make(chan string, 10)} }
