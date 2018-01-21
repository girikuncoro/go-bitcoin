package vip_client

func (c *Client) GetInfo() ([]byte, error) {
	params := c.encodedParams(ParamGetInfo)
	req, err := c.setRequest("POST", BitcoinURL, params)
	if err != nil {
		return nil, err
	}
	c.setPrivateHeader(req, params)

	res, err := c.response(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
