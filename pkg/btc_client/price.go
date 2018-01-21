package btc_client

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

func (c *Client) GetPrice(coinName string) (*Ticker, error) {
	req, err := c.setRequest("GET", c.buildTickerURL(coinName), "")
	if err != nil {
		return nil, err
	}

	res, err := c.response(req)
	if err != nil {
		return nil, err
	}

	t := &Ticker{}
	err = json.Unmarshal(res, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) buildTickerURL(coinName string) string {
	url := fmt.Sprintf(TickerURL, coinName)
	return url
}
