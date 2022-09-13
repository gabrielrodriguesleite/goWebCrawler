package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Web Crawler Go v1.0.0")

	url := "https://aprendagolang.com.br"

	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Não foi possível acessar: %s\nErro: %v", url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("[Erro] Status code diferente de 200: %v", resp.StatusCode))
	}

	// fmt.Println("Body:", resp.Body)
	body, err := html.Parse(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Erro ao traduzir html: %v", err))
	}

	// fmt.Println("Body: ", body)
	linkVisited(body)
}

func linkVisited(body *html.Node) {
	if body.Type == html.ElementNode && body.Data == "a" {
		fmt.Println(body.Data)
	}

	for c := body.FirstChild; c != nil; c = c.NextSibling {
		linkVisited(c)
	}
}
