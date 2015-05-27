package maicoin

import (
	"encoding/json"
	"strconv"
)

func (param *CheckoutParam) Build() string {
	if param.Checkout.Items == nil {
		param.Checkout.Items = make([]CheckoutParamItem, 0, 1)
	}
	v2b, _ := json.Marshal(param)
	return string(v2b)
}

func (param *CheckoutParam) SetCheckoutData(amount float64, currency string,
	returnUrl string, cancelUrl string, callbackUrl string,
	merchantRefId string, posData string, locale string) {
	param.Checkout.Amount = strconv.FormatFloat(amount, 'f', 8, 64)
	param.Checkout.Currency = currency
	param.Checkout.ReturnUrl = returnUrl
	param.Checkout.CancelUrl = cancelUrl
	param.Checkout.CallbackUrl = callbackUrl
	param.Checkout.MerchantRefId = merchantRefId
	param.Checkout.PosData = posData
	param.Checkout.Locale = locale
}

func (param *CheckoutParam) SetBuyerData(name string, address1 string,
	address2 string, city string, state string, zip string,
	email string, phone string, country string) {
	param.Checkout.Buyer.BuyerName = name
	param.Checkout.Buyer.BuyerAddress1 = address1
	param.Checkout.Buyer.BuyerAddress2 = address2
	param.Checkout.Buyer.BuyerCity = city
	param.Checkout.Buyer.BuyerState = state
	param.Checkout.Buyer.BuyerZip = zip
	param.Checkout.Buyer.BuyerEmail = email
	param.Checkout.Buyer.BuyerPhone = phone
	param.Checkout.Buyer.BuyerCountry = country
}

func (param *CheckoutParam) AddItem(desc string, code string, price string,
	currency string, isPhysical string) {
	var item = CheckoutParamItem{}
	item.Item.Description = desc
	item.Item.Code = code
	item.Item.Price = price
	item.Item.Currency = currency
	item.Item.IsPhysical = isPhysical
	if param.Checkout.Items == nil {
		param.Checkout.Items = make([]CheckoutParamItem, 0, 1)
	}
	param.Checkout.Items = append(param.Checkout.Items, item)
}

func (c *Client) CreateCheckout(jsonForm string) (Checkout, error) {
	body, err := c.HttpVerb(HttpPost, "/checkouts", nil, jsonForm)
	var response Checkout
	err = json.Unmarshal(body, &response)
	return response, err
}
