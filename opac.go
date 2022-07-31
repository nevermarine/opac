package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"golang.org/x/net/html"
)

const (
	library_link = "http://opac.hse.ru/absopac/app/webroot/index.php?url=/user_card"
	root_url     = "http://opac.hse.ru"
)

var (
	login, password string
	login_field     = "data[User][CodbarU]"
	password_field  = "data[User][MotPasse]"
	dry_run, help   bool
	err             error
)

func main() {
	ArgsInit()

	// just exit
	if login == "" || password == "" || help {
		flag.PrintDefaults()
		return
	}

	data := url.Values{
		login_field:    {login},
		password_field: {password},
	}
	doc, err := DataRequest(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

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
				fmt.Fprintln(os.Stderr, "Error:", err)
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

func ArgsInit() {
	flag.StringVarP(&login, "login", "l", "", "User login for HSE OPAC. Required.")
	flag.StringVarP(&password, "pass", "p", "", "Password for HSE OPAC. Required.")
	flag.BoolVarP(&dry_run, "dry-run", "n", false, "Do nothing, just print to stdout.")
	flag.BoolVarP(&help, "help", "h", false, "Print this help message.")
	flag.Parse()
}

func DataRequest(data url.Values) (*html.Node, error) {
	resp, err := http.PostForm(library_link, data)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return doc, nil
}
