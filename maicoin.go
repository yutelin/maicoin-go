package maicoin

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
	"time"
	//"encoding/json"
	"strconv"
)

// Constants
const (
	MAICOIN_API_ENDPOINT = "https://api.maicoin.com/v1/"
)

type Verb int

const (
	HttpGet    = iota
	HttpPost   = iota
	HttpPut    = iota
	HttpDelete = iota
)

// Client

type Client struct {
	ApiKey     string
	ApiSecret  string
	httpClient *http.Client
}

// ComputeHmac256
func ComputeHmac256(secret string, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) HttpVerb(verb Verb, path string, params map[string]interface{}) ([]byte, error) {
	// Build HTTP client
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	apiURL := MAICOIN_API_ENDPOINT + path

	//Headers
	nonce := strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	hmacMessage := nonce + apiURL
	if params == nil {
		params = make(map[string]interface{})
	}
	if verb == HttpPost || verb == HttpPut {
		postBody, _ := json.Marshal(params)
		hmacMessage = hmacMessage + string(postBody)
	}
	signature := ComputeHmac256(c.ApiSecret, hmacMessage)
	// fmt.Println("nonce:", nonce)
	// fmt.Println("hmacMessage", hmacMessage)
	// fmt.Println("signature", signature)

	apiURL = apiURL //+ "/?" + params.Encode()
	var req *http.Request
	var err error
	switch verb {
	case HttpGet:
		req, err = http.NewRequest("GET", apiURL, nil)
	case HttpPost:
		postBody, _ := json.Marshal(params)
		req, err = http.NewRequest("POST", apiURL, bytes.NewReader(postBody))
	case HttpPut:
		postBody, _ := json.Marshal(params)
		req, err = http.NewRequest("PUT", apiURL, bytes.NewReader(postBody))
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("ACCESS_KEY", c.ApiKey)
	req.Header.Set("ACCESS_SIGNATURE", signature)
	req.Header.Set("ACCESS_NONCE", nonce)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	return c.makeRequest(req)
}

func (c *Client) makeRequest(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Make sure we close the body stream no matter what
	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	// Check status code
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		//return nil, fmt.Errorf("Invalid HTTP response code: %d", resp.StatusCode)
	}

	// Return
	return body, err
}
