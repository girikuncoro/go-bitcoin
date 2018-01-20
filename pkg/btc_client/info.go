package btc_client

func (c *Client) GetInfo() (string, error) {
	params := c.encodedParams(ParamGetInfo)
	req, err := c.setRequest("POST", params)
	if err != nil {
		return "", err
	}

	res, err := c.response(req)
	if err != nil {
		return "", err
	}

	return res, nil
}
