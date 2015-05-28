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
	"strconv"
	"strings"
	"time"
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
	ApiKey      string
	ApiSecret   string
	AccessToken string
	httpClient  *http.Client
}

// ComputeHmac256
func ComputeHmac256(secret string, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) HttpVerb(verb Verb, path string, params map[string]interface{}, jsonForm string) ([]byte, error) {
	// Build HTTP client
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	apiURL := MAICOIN_API_ENDPOINT + path
	var req *http.Request
	var err error
	switch verb {
	case HttpGet:
		req, err = http.NewRequest("GET", apiURL, nil)
	case HttpDelete:
		req, err = http.NewRequest("DELETE", apiURL, nil)
	case HttpPost:
		if len(jsonForm) > 0 {
			req, err = http.NewRequest("POST", apiURL, strings.NewReader(jsonForm))
		} else {
			postBody, _ := json.Marshal(params)
			req, err = http.NewRequest("POST", apiURL, bytes.NewReader(postBody))
		}
	case HttpPut:
		if len(jsonForm) > 0 {
			req, err = http.NewRequest("PUT", apiURL, strings.NewReader(jsonForm))
		} else {
			postBody, _ := json.Marshal(params)
			req, err = http.NewRequest("PUT", apiURL, bytes.NewReader(postBody))
		}
	}
	if err != nil {
		return nil, err
	}

	if len(c.ApiKey) > 0 && len(c.ApiSecret) > 0 {
		//Headers
		nonce := strconv.FormatInt(time.Now().UnixNano()/1000, 10)
		hmacMessage := nonce + apiURL
		if params == nil {
			params = make(map[string]interface{})
		}
		if verb == HttpPost || verb == HttpPut {
			if len(jsonForm) > 0 {
				hmacMessage = hmacMessage + jsonForm
			} else {
				postBody, _ := json.Marshal(params)
				hmacMessage = hmacMessage + string(postBody)
			}
		}
		signature := ComputeHmac256(c.ApiSecret, hmacMessage)
		req.Header.Set("ACCESS_KEY", c.ApiKey)
		req.Header.Set("ACCESS_SIGNATURE", signature)
		req.Header.Set("ACCESS_NONCE", nonce)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("AUTHORIZATION", "Bearer "+c.AccessToken)
	}

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

	// Check status code
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
	}

	// Return
	return body, err
}
