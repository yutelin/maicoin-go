package maicoin

import (
	"encoding/json"
	// "fmt"
)

//Balance
func (c *Client) Balance() (Balance, error) {
	body, err := c.HttpVerb(HttpGet, "/account/balance", nil)
	var response Balance
	err = json.Unmarshal(body, &response)
	return response, err
}

// User
func (c *Client) User() (User, error) {
	body, err := c.HttpVerb(HttpGet, "/user", nil)
	var response User
	err = json.Unmarshal(body, &response)
	return response, err
}

// Receive address
func (c *Client) ReceiveAddress(currency string) (Address, error) {
	body, err := c.HttpVerb(HttpGet, "/account/receive_address/"+currency, nil)
	var response Address
	err = json.Unmarshal(body, &response)
	return response, err
}

// Receive address
func (c *Client) Addresses(currency string) (Addresses, error) {
	body, err := c.HttpVerb(HttpGet, "/account/addresses/"+currency, nil)
	var response Addresses
	err = json.Unmarshal(body, &response)
	return response, err
}

// Generate Receive Address
func (c *Client) GenerateReceiveAddress(currency string) (Address, error) {
	params := make(map[string]interface{})
	params["currency"] = currency
	body, err := c.HttpVerb(HttpPost, "/account/receive_address", params)
	var response Address
	err = json.Unmarshal(body, &response)
	return response, err
}

// Create account pin
func (c *Client) CreateAccountPin(pin string) (Result, error) {
	params := make(map[string]interface{})
	params["pin"] = pin
	body, err := c.HttpVerb(HttpPost, "/user/account_pin", params)
	var response Result
	err = json.Unmarshal(body, &response)
	return response, err
}

// Update account pin
func (c *Client) UpdateAccountPin(oldPin string, newPin string) (Result, error) {
	params := make(map[string]interface{})
	params["old_pin"] = oldPin
	params["new_pin"] = newPin
	body, err := c.HttpVerb(HttpPut, "/user/account_pin", params)
	var response Result
	err = json.Unmarshal(body, &response)
	return response, err
}
