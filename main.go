package main

import (
	"flag"
	"fmt"

	"github.com/gabrielrodriguesleite/goWebCrawler/crawler"
	"github.com/gabrielrodriguesleite/goWebCrawler/website"
)

var (
	link   string
	action string
)

func init() {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br", "url para iniciar visitas")
	flag.StringVar(&action, "action", "website", "qual serviço usar")
}

func main() {
	flag.Parse()
	fmt.Println("Web Crawler Go v1.0.0")

	switch action {
	case "website":
		website.Run()
	case "webcrawler":
		done := make(chan bool)
		go crawler.VisitLink(link)
		<-done
	default:
		fmt.Printf("Action não reconhecida: %s\n", action)
	}
}
