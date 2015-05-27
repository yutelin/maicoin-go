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

type AddressObject struct {
	Address struct {
		Address   string
		CreatedAt string `json:"created_at"`
	}
}

type Addresses struct {
	Success      bool
	Code         int
	Errors       []string
	Addresses    []AddressObject
	CurrencyType string `json:"currency_type"`
}

type MaicoinBankInformation struct {
	MaicoinBankCode          string `json:"maicoin_bank_code"`
	MaicoinBankAccountName   string `json:"maicoin_bank_account_name"`
	MaicoinBankAccountNumber string `json:"maicoin_bank_account_number"`
}

type Order struct {
	Type                   string
	Status                 string
	Txid                   string
	CreatedAt              string `json:"created_at"`
	Notes                  string
	CoinAmount             string `json:"coin_amount"`
	CoinCurrency           string `json:"coin_currency"`
	CoinPrice              string `json:"coin_price"`
	TotalCost              string `json:"total_cost"`
	Currency               string
	MaicoinBankInformation MaicoinBankInformation `json:"maicoin_bank_information"`
}

type OrderResponse struct {
	Success bool
	Code    int
	Errors  []string
	Order   Order
}

type OrderObject struct {
	Order Order
}

type Orders struct {
	Success     bool
	Code        int
	Errors      []string
	Orders      []OrderObject
	CurrentPage int `json:"current_page"`
	Count       int
	NumOfPages  int `json:"num_of_pages"`
}

type Transaction struct {
	Id              string
	CreatedAt       string `json:"created_at"`
	Amount          string
	Currency        string
	Status          string
	TransactionType string `json:"transaction_type"`
	Notes           string
	Sender          struct {
		Name  string
		Email string
	}
	Recipient struct {
		Name  string
		Email string
	}
}

type TransactionObject struct {
	Transaction Transaction
}

type Transactions struct {
	Success      bool
	Code         int
	Errors       []string
	Transactions []TransactionObject
	CurrentPage  int `json:"current_page"`
	Count        int
	NumOfPages   int `json:"num_of_pages"`
	CurrentUser  struct {
		Name  string
		Email string
	} `json:"current_user"`
	Balance struct {
		Amount   string
		Currency string
	}
	NativeBalance struct {
		Amount   string
		Currency string
	} `json:"native_balance"`
}

type TransactionResponse struct {
	Success     bool
	Code        int
	Errors      []string
	Transaction Transaction
}
