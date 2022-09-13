package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Web Crawler Go v1.0.0")

	url := "https://aprendagolang.com.br"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("Status code diferente de 200")
	}

	fmt.Println("Body:", resp.Body)
}
