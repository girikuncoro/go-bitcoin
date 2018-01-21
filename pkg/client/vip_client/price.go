package vip_client

import (
	"encoding/json"
	"fmt"
)

type Ticker struct {
	Ticker struct {
		High string `json:"high"`
		Low  string `json:"low"`
		Last string `json:"last"`
	}
}

func (c *VipClient) GetPrice(coinName string) (string, error) {
	req, err := c.baseClient.Request("GET", c.buildTickerURL(coinName), "")
	if err != nil {
		return "", err
	}

	res, err := c.baseClient.Response(req)
	if err != nil {
		return "", err
	}

	t := &Ticker{}
	err = json.Unmarshal(res, &t)
	if err != nil {
		return "", err
	}

	return t.Ticker.Last, nil
}

func (c *VipClient) buildTickerURL(coinName string) string {
	url := fmt.Sprintf(TickerURL, coinName)
	return url
}
