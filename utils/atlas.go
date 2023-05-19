package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	baseUrl = "http://hadoop102:21000/api/"
)

func Call(part, username, password, method string, queryParams map[string]string, body interface{}) ([]byte, error) {
	url := baseUrl + part
	req := &http.Request{}
	client := &http.Client{}
	switch method {
	case "POST":
		req, _ = post(url, body, queryParams, nil)
	case "GET":
		req, _ = get(url, queryParams, nil)
	}
	req.SetBasicAuth(username, password)
	resp, _ := client.Do(req)
	s, _ := ioutil.ReadAll(resp.Body)
	return s, nil
}

func get(url string, params map[string]string, headers map[string]string) (*http.Request, error) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	log.Printf("Go %s URL : %s \n", http.MethodGet, req.URL.String())
	return req, nil
}
func post(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Request, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	log.Printf("Go %s URL : %s \n", http.MethodPost, req.URL.String())
	return req, nil
}
