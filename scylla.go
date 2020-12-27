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
	}
}

/*
	@query Your query using the Lucene query syntax
	@size  how many items you want returned
	@start Record number to start from.
 */
func Query(query string, size int, start int) ([]Result, error){
		// Currently having TLS issues using the IP, and HTTP via IP redirects to domain
		transportConfig := &http.Transport{
			TLSClientConfig:  &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: transportConfig}

		//url := "https://scylla.sh/search"
		url := "https://44.235.17.188/search"
		payload := fmt.Sprintf("%s?q=%s&size=%d&start=%d",url, query, size, start)

		//fmt.Println(payload)
		res, err := client.Get(payload)
		defer res.Body.Close()

		if err != nil{
			errors.New(err.Error())
		}

		if res.StatusCode != http.StatusOK{
				errors.New("returned a non OK header")
		}

		var result []Result
		err = json.NewDecoder(res.Body).Decode(&result)
		//err := json.Unmarshal(data, &result)
		return result, err
}
/**
// Simple usuage example.
func main(){
	test, err := Query("username:jb", 10, 0)

	if err != nil{
		log.Fatal(err)
	}

	for _, it := range test {
		fmt.Printf("%#v\n", it)
	}

}
 */