package maicoin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (c *Client) Transactions(page int, limit int) (Transactions, error) {
	params := url.Values{}
	params.Set("page", strconv.Itoa(page))
	params.Set("limit", strconv.Itoa(limit))
	body, err := c.HttpVerb(HttpGet, "/transactions?"+params.Encode(), nil, "")
	var response Transactions
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) Transaction(txid string) (TransactionResponse, error) {
	body, err := c.HttpVerb(HttpGet, "/transactions/"+txid, nil, "")
	var response TransactionResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) RequestTransaction(address string, amount float64, currency string) (TransactionResponse, error) {
	params := make(map[string]interface{})
	params["type"] = "request"
	params["currency"] = currency
	params["address"] = address
	params["amount"] = strconv.FormatFloat(amount, 'f', 8, 64)
	body, err := c.HttpVerb(HttpPost, "/transactions", params, "")
	var response TransactionResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) CancelRequestTransaction(txid string) (TransactionResponse, error) {
	body, err := c.HttpVerb(HttpDelete, "/transactions/"+txid, nil, "")
	var response TransactionResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) ApproveRequestTransaction(txid string, accountPin string) (TransactionResponse, error) {
	params := make(map[string]interface{})
	params["account_pin"] = accountPin
	body, err := c.HttpVerb(HttpPut, "/transactions/"+txid+"/approve", params, "")
	var response TransactionResponse
	err = json.Unmarshal(body, &response)
	return response, err
}
