package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please specify address")
		return
	}
	addr := args[1]

	parsedUrl, err := url.Parse(addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{Transport: transport}

	resp, err := client.Get(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	splitted := strings.Split(parsedUrl.Path, "/")
	fileName := "page"
	if l := len(splitted); l != 0 {
		fileName = splitted[l-1]
	}
	if !strings.Contains(fileName, ".") {
		fileName += ".html"
	}
	os.WriteFile(fileName, body, 0600)
}
