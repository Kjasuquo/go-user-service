package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Service struct {
	endpoint string
}

func CreatService(endpoint string) Service {
	return Service{endpoint: endpoint}
}

//const endpoint = "http://host.docker.internal:8001/api/"

func (service *Service) Get(path, cookie string) (*http.Response, error) {
	return service.request("GET", path, cookie, nil)
}

func (service *Service) Post(path, cookie string, body map[string]string) (*http.Response, error) {
	return service.request("POST", path, cookie, body)
}

func (service *Service) Put(path, cookie string, body map[string]string) (*http.Response, error) {
	return service.request("PUT", path, cookie, body)
}

func (service *Service) Delete(path, cookie string, body map[string]string) (*http.Response, error) {
	return service.request("DELETE", path, cookie, body)
}

func (service *Service) request(method, path, cookie string, body map[string]string) (*http.Response, error) {
	var data io.Reader = nil

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			log.Fatalf("Can't marshal body into json: %v\n", err)
		}

		data = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, service.endpoint+path, data)
	if err != nil {
		log.Fatalf("error in NewRequest: %v\n", err)
	}

	req.Header.Add("Content-Type", "application/json")

	if cookie != "" {
		req.Header.Add("Cookie", "jwt="+cookie)
	}

	client := &http.Client{}
	return client.Do(req)

}
