package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	nrHttp "github.com/newrelic/infra-integrations-sdk/http"
	"github.com/newrelic/nri-f5/src/arguments"
)

// F5Client represents a client that is able to make requests to the F5 iControl API.
type F5Client struct {
	HTTPClient         *http.Client
	Username           string
	Password           string
	AuthToken          string
	BaseURL            string
	LoginProviderName  string
	AuthURL            string
	AuthPath           string
	LoginReferenceLink string
}

//const loginEndpoint = "/mgmt/shared/authn/login"

// NewClient takes in arguments and creates and returns a client that will talk to the F5 API, or an error if one cannot be created
func NewClient(args *arguments.ArgumentList) (*F5Client, error) {
	httpClient, err := nrHttp.New(args.CABundleFile, args.CABundleDir, time.Duration(args.Timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	return &F5Client{
		HTTPClient:         httpClient,
		Username:           args.Username,
		Password:           args.Password,
		AuthToken:          "",
		BaseURL:            "https://" + args.Hostname + ":" + strconv.Itoa(args.Port),
		AuthURL:            "https://" + args.AuthHost + ":" + strconv.Itoa(args.AuthPort),
		LoginProviderName:  args.LoginProviderName,
		AuthPath:           args.AuthPath,
		LoginReferenceLink: args.LoginReferenceLink,
	}, nil
}

// Request is a shortcut for making a GET request without a request body
func (c *F5Client) Request(endpoint string, model interface{}) error {
	return c.DoRequest(http.MethodGet, endpoint, "", model)
}

// DoRequest makes a request to the given endpoint using the given request body, storing the result in the model if possible.
// An error is returned if either step cannot be completed.
func (c *F5Client) DoRequest(method, endpoint, body string, model interface{}) error {
	var req *http.Request
	var err error

	if endpoint == c.AuthPath {
		req, err = http.NewRequest(method, c.AuthURL+endpoint, strings.NewReader(body))
	} else {
		if c.AuthToken == "" {
			return fmt.Errorf("client is not logged in")
		}
		req, err = http.NewRequest(method, c.BaseURL+endpoint, strings.NewReader(body))
		req.Header.Add("X-F5-Auth-Token", c.AuthToken)
	}

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if err = checkStatusCode(res); err != nil {
		return fmt.Errorf("request failed for endpoint %s: %s", endpoint, err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(model)
	if err != nil {
		return err
	}

	return nil
}

// LogIn attempts to retrieve an auth token from the API using the credentials the client was created with and returns nil.
// If login is unsuccessful an error is returned
func (c *F5Client) LogIn() error {
	loginArgs := map[string]interface{}{
		"loginProviderName": c.LoginProviderName,
		"username":          c.Username,
		"password":          c.Password,
	}

	if c.LoginReferenceLink != "" {
		lr := map[string]string{
			"link": c.LoginReferenceLink,
		}
		loginArgs["loginReference"] = lr
	}
	loginBody, err := json.Marshal(loginArgs)
	if err != nil {
		return err
	}

	var loginResponse tokenResponse
	err = c.DoRequest(http.MethodPost, c.AuthPath, string(loginBody), &loginResponse)
	if err != nil {
		return err
	}

	if loginResponse.Token.Token == nil {
		return fmt.Errorf("couldn't get auth token from response")
	}

	// successful request, extract token
	c.AuthToken = *loginResponse.Token.Token
	return nil
}

func checkStatusCode(response *http.Response) error {
	if response.StatusCode != 200 {
		return fmt.Errorf("response contained non-200 status code %d", response.StatusCode)
	}
	return nil
}

type tokenResponse struct {
	Token *tokenStruct `json:"token"`
}

type tokenStruct struct {
	Token *string `json:"token"`
}
