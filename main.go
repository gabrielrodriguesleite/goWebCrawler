package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var links []string

func main() {
	fmt.Println("Web Crawler Go v1.0.0")

	visitLink("https://aprendagolang.com.br")

	fmt.Println("Qtd links: ", len(links))
}

func visitLink(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Não foi possível acessar: %s\nErro: %v", url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("[Erro] Status code diferente de 200: %v", resp.StatusCode))
	}

	// fmt.Println("Body:", resp.Body)
	element, err := html.Parse(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Erro ao traduzir html: %v", err))
	}

	// fmt.Println("Body: ", body)
	extractLinks(element)
}

func extractLinks(element *html.Node) {
	if element.Type == html.ElementNode && element.Data == "a" {
		for _, attr := range element.Attr {
			if attr.Key != "href" {
				continue // continua buscando até achar href
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}

			links = append(links, link.String())

			visitLink(link.String())
			// fmt.Println(attr.Val)
		}
	}

	for c := element.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}
