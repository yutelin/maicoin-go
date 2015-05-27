package maicoin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (c *Client) BuyOrder(amount float64, currency string) (OrderResponse, error) {
	params := make(map[string]interface{})
	params["currency"] = currency
	params["type"] = "buy"
	params["amount"] = strconv.FormatFloat(amount, 'f', 8, 64)
	body, err := c.HttpVerb(HttpPost, "/orders", params, "")
	var response OrderResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) SellOrder(amount float64, currency string) (OrderResponse, error) {
	params := make(map[string]interface{})
	params["currency"] = currency
	params["type"] = "sell"
	params["amount"] = strconv.FormatFloat(amount, 'f', 8, 64)
	body, err := c.HttpVerb(HttpPost, "/orders", params, "")
	var response OrderResponse
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) Orders(page int, limit int) (Orders, error) {
	params := url.Values{}
	params.Set("page", strconv.Itoa(page))
	params.Set("limit", strconv.Itoa(limit))
	body, err := c.HttpVerb(HttpGet, "/orders?"+params.Encode(), nil, "")
	var response Orders
	err = json.Unmarshal(body, &response)
	return response, err
}

func (c *Client) Order(txid string) (OrderResponse, error) {
	body, err := c.HttpVerb(HttpGet, "/orders/"+txid, nil, "")
	var response OrderResponse
	err = json.Unmarshal(body, &response)
	return response, err
}
