package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please specify adresse")
		return
	}
	addr := args[1]

	parsedUrl, err := url.Parse(addr)
	if err != nil {
		fmt.Println("The following error occured", err)
		return
	}

	splitted := strings.Split(parsedUrl.Path, "/")
	var pageName string
	if l := len(splitted); l > 0 {
		pageName = splitted[l-1]
	} else {
		pageName = "page"
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}

	resp, err := client.Get(addr)
	if err != nil {
		fmt.Println("The following error occured", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("The following error occured", err)
		return
	}

	fileName := pageName
	if !strings.Contains(fileName, ".") {
		fileName += ".html"
	}

	ioutil.WriteFile(fileName, body, 0600)
}
