package vip_client

func (c *VipClient) GetInfo() ([]byte, error) {
	params := c.encodedParams(ParamGetInfo)
	req, err := c.baseClient.Request("POST", BitcoinURL, params)
	if err != nil {
		return nil, err
	}
	c.setPrivateHeader(req, params)

	res, err := c.baseClient.Response(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
