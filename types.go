package maicoin

type Result struct {
	Success bool
	Code    int
	Errors  []string
}

type Price struct {
	Success   bool
	Code      int
	Errors    []string
	SellPrice string `json:"sell_price"`
	BuyPrice  string `json:"buy_price"`
	Price     string
	Currency  string
}

type Currencies struct {
	Success    bool
	Code       int
	Errors     []string
	Currencies []string
}

type User struct {
	Success bool
	Code    int
	Errors  []string
	User    struct {
		Email         string
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		Locale        string `json:"locale"`
		AccountPinSet bool   `json:"account_pin_set"`
		PhoneSet      bool   `json:"phone_set"`
	}
}

type Balance struct {
	Success      bool
	Code         int
	Errors       []string
	CoinAmount   string `json:"coin_amount"`
	CoinCurrency string `json:"coin_currency"`
}

type Address struct {
	Success      bool
	Code         int
	Errors       []string
	Address      string
	CurrencyType string `json:"currency_type"`
}

type AddressType struct {
	Address struct {
		Address   string
		CreatedAt string `json:"created_at"`
	}
}

type Addresses struct {
	Success      bool
	Code         int
	Errors       []string
	Addresses    []AddressType
	CurrencyType string `json:"currency_type"`
}
