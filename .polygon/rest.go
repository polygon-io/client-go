package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected a URI for fetching the REST spec")
	}
	uri := os.Args[1]

	var body []byte
	if isURL(uri) {
		resp, err := http.Get(uri)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("expected arg to be a valid URI")
	}

	var out bytes.Buffer
	err := json.Indent(&out, body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(".polygon/rest.json", out.Bytes(), 0644)
}

func isURL(uri string) bool {
	u, err := url.Parse(uri)
	return err == nil && u.Scheme != "" && u.Host != ""
}
