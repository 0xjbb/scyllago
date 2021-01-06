package scyllago

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Result struct {
	Id     string
	Fields struct {
		Username string
		Password string
		Email    string
		Domain   string
		Ip       string
		Passhash string
	}
}

/*
	@query Your query using the Lucene query syntax
	@size  how many items you want returned
	@start Record number to start from.
*/
func Query(query string, size int, start int) ([]Result, error) {

	if query == "" {
		return nil, errors.New("please enter a lucene query string")
	}

	if size == 0 {
		return nil, errors.New("size is 0") // why make unnecessary API calls.
	}

	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: transportConfig}

	url := "https://scylla.sh/search" // domain is fixed.
	payload := fmt.Sprintf("%s?q=%s&size=%d&start=%d", url, query, size, start)

	res, err := client.Get(payload)
	defer res.Body.Close()

	if err != nil {
		errors.New(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		// maybe just return the response body as the error
		// @TODO change to return the API's error message.
		errors.New("returned a non OK header, check your query")
	}

	var result []Result
	err = json.NewDecoder(res.Body).Decode(&result)

	return result, err
}