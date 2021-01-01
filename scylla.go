package scyllago

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)
/**
	@TODO Maybe change from struct to map, idk I think maps are faster?
 */
type Result struct {
	Id     string
	Fields struct {
		Username string
		Password string
		Email    string
		Domain   string
	}
}

/*
	@query Your query using the Lucene query syntax
	@size  how many items you want returned
	@start Record number to start from.
 */
func Query(query string, size int, start int) ([]Result, error){
		// Currently having TLS issues using the IP, and HTTP via IP redirects to domain

		// @TODO test for stuff that don't work :kek:
		if query == ""{
			return nil, errors.New("please enter a Lucene query string")
		}

		if size == 0 {
			return nil, errors.New("size is 0")// why make unnecessary API calls.
		}

		transportConfig := &http.Transport{
			TLSClientConfig:  &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: transportConfig}

		//url := "https://44.235.17.188/search" // remove when domain is fixed.
		url := "https://scylla.sh/search" // domain is fixed.
		payload := fmt.Sprintf("%s?q=%s&size=%d&start=%d",url, query, size, start)

		//fmt.Println(payload)
		res, err := client.Get(payload)
		defer res.Body.Close()

		if err != nil{
			errors.New(err.Error())
		}

		if res.StatusCode != http.StatusOK{
			// maybe just return the response body as the error
				errors.New("returned a non OK header, check your query")
		}

		var result []Result
		err = json.NewDecoder(res.Body).Decode(&result)
		//err := json.Unmarshal(data, &result)
		return result, err
}