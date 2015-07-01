// Main contains a Web server that serves as a proxy for a Drupal web site
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {

	var url string
	url = "http://dev-go.pantheon.io/node/" + req.URL.Path[1:] + "?_format=json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body_text, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	rw.Write(body_text)
}
