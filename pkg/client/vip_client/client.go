package vip_client

import (
	"net/http"
	"net/url"

	"github.com/girikuncoro/go-bitcoin/pkg/client/base_client"
)

type VipClient struct {
	auth       *Auth
	baseClient *base_client.BaseClient
}

func NewVipClient(baseClient *base_client.BaseClient) *VipClient {
	return &VipClient{
		baseClient: baseClient,
	}
}

func (c *VipClient) encodedParams(method string) string {
	if method == "" {
		return ""
	}

	v := url.Values{}
	v.Set(ParamNonce, c.auth.Nonce())
	v.Set(ParamMethod, method)
	return v.Encode()
}

func (c *VipClient) setPrivateHeader(req *http.Request, encodedParams string) *http.Request {
	req.Header.Add(HeaderContentType, HeaderFormUrlEncoded)
	req.Header.Add(HeaderSignedMessage, c.auth.Sign(encodedParams))
	req.Header.Add(HeaderApiKey, c.auth.ApiKey)

	return req
}
