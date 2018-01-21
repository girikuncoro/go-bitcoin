package base_client

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type BaseClient struct {
	http *http.Client
}

func NewBaseClient() *BaseClient {
	return &BaseClient{
		http: &http.Client{},
	}
}

func (c *BaseClient) Request(reqMethod, url, encodedParams string) (*http.Request, error) {
	req, err := http.NewRequest(reqMethod, url, strings.NewReader(encodedParams))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *BaseClient) Response(req *http.Request) ([]byte, error) {
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
