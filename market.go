package maicoin

import (
	"encoding/json"
)

// Prices
func (c *Client) Prices(currency string) (Price, error) {
	body, err := c.HttpVerb(HttpGet, "/prices/"+currency, nil)
	var response Price
	err = json.Unmarshal(body, &response)
	return response, err
}

// Currencies
func (c *Client) Currencies() (Currencies, error) {
	body, err := c.HttpVerb(HttpGet, "/currencies", nil)
	var response Currencies
	err = json.Unmarshal(body, &response)
	return response, err
}
