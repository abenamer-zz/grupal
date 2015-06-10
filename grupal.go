package main // grupal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://dev-d8go.pantheon.io/d8go/node/1", nil)
	req.Header.Set("Accept-Encoding:", "gzip, deflate, sdch")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
