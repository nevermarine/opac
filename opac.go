package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	flag "github.com/spf13/pflag"
	"golang.org/x/net/html"
)

const (
	library_link = "http://opac.hse.ru/absopac/app/webroot/index.php?url=/user_card"
	root_url     = "http://opac.hse.ru"
)

func main() {
	var login, password string
	var dry_run, help bool
	var err error

	flag.StringVarP(&login, "login", "l", "", "User login for HSE OPAC. Required.")
	flag.StringVarP(&password, "pass", "p", "", "Password for HSE OPAC. Required.")
	flag.BoolVarP(&dry_run, "dry-run", "n", false, "Do nothing, just print to stdout.")
	flag.BoolVarP(&help, "help", "h", false, "Print this help message.")
	flag.Parse()

	if login == "" || password == "" || help {
		flag.PrintDefaults()
		return
	}

	data := url.Values{
		"data[User][CodbarU]":  {login},
		"data[User][MotPasse]": {password},
	}

	resp, _ := http.PostForm(library_link, data)
	doc, _ := html.Parse(resp.Body)
	resp.Body.Close()
	links := make(chan string)
	go func() {
		TraverseFind(doc, links)
		close(links)
	}()

	for {
		v, ok := <-links
		if !ok {
			break
		}
		if !dry_run {
			_, err = http.PostForm(root_url+v, data)
			if err != nil {
				fmt.Println("Failed to renew the book.")
			} else {
				fmt.Println("Successfully renewed the book.")
			}
		} else {
			fmt.Printf("Renewal link: %s\n", root_url+v)
		}
	}

}

func TraverseFind(n *html.Node, ch chan string) {
	if strings.Contains(n.Data, "Продлить") {
		// parent contains the renewal link
		for _, i := range n.Parent.Attr {
			ch <- i.Val
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TraverseFind(c, ch)
	}
}
