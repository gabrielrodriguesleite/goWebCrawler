package website

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/gabrielrodriguesleite/goWebCrawler/crawler"
	// "github.com/gorilla/websocket"
	"nhooyr.io/websocket"
)

func websocketHandle() func(http.ResponseWriter, *http.Request) {
	tmpl, err := template.ParseFiles("website/templates/websocket.html")
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		website := r.FormValue("website")
		if website == "" {
			http.Error(w, "O campo website não pode estar vazio", http.StatusBadRequest)
			return
		}

		wc := crawler.New()
		go wc.VisitLink(website)

		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}

		subcriber(r.Context(), c, wc.Log())
		tmpl.Execute(w, nil)
	}

}

func subcriber(ctx context.Context, c *websocket.Conn, logs <-chan string) error {
	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-logs:
			err := writeTimeout(ctx, c, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func writeTimeout(ctx context.Context, c *websocket.Conn, msg string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	return c.Write(ctx, websocket.MessageText, []byte(msg))
}
