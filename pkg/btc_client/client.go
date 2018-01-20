package btc_client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Clienter interface {
	GetInfo()
}

type Client struct {
	auth *Auth
	http *http.Client
	req  *http.Request
}

func NewClient(auth *Auth, http *http.Client) *Client {
	return &Client{
		auth: auth,
		http: http,
	}
}

func (c *Client) encodedParams(method string) string {
	if method == "" {
		return ""
	}

	v := url.Values{}
	v.Set(ParamNonce, c.auth.Nonce())
	v.Set(ParamMethod, method)
	return v.Encode()
}

func (c *Client) setRequest(reqMethod, encodedParams string) (*http.Request, error) {
	req, err := http.NewRequest(reqMethod, BitcoinURL, strings.NewReader(encodedParams))
	if err != nil {
		return nil, err
	}
	req.Header.Add(HeaderContentType, HeaderFormUrlEncoded)
	req.Header.Add(HeaderSignedMessage, c.auth.Sign(encodedParams))
	req.Header.Add(HeaderApiKey, c.auth.ApiKey)

	return req, nil
}

func (c *Client) response(req *http.Request) (string, error) {
	res, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
