package gemini_client

import (
	"encoding/json"
	"fmt"
)

type Ticker struct {
	Bid  string `json:"bid"`
	Ask  string `json:"ask"`
	Last string `json:"last"`
}

func (c *GeminiClient) GetPrice(coinName string) (string, error) {
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

	return t.Last, nil
}

func (c *GeminiClient) buildTickerURL(coinName string) string {
	url := fmt.Sprintf(TickerURL, coinName)
	return url
}
