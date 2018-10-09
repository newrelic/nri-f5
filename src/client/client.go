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
	httpClient *http.Client
	username   string
	password   string
	authToken  string
	baseURL    string
}

const loginEndpoint = "/mgmt/shared/authn/login"

// NewClient takes in arguments and creates and returns a client that will talk to the F5 API, or an error if one cannot be created
func NewClient(args *arguments.ArgumentList) (*F5Client, error) {
	httpClient, err := nrHttp.New(args.CABundleFile, args.CABundleDir, time.Duration(args.Timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	return &F5Client{
		httpClient: httpClient,
		username:   args.Username,
		password:   args.Password,
		authToken:  "",
		baseURL:    "https://" + args.Hostname + ":" + strconv.Itoa(args.Port),
	}, nil
}

// Request is a shortcut for making a GET request without a request body
func (c *F5Client) Request(endpoint string, model interface{}) error {
	return c.DoRequest(http.MethodGet, endpoint, "", model)
}

// DoRequest makes a request to the given endpoint using the given request body, storing the result in the model if possible.
// An error is returned if either step cannot be completed.
func (c *F5Client) DoRequest(method, endpoint, body string, model interface{}) error {
	req, err := http.NewRequest(method, c.baseURL+endpoint, strings.NewReader(body))
	if err != nil {
		return err
	}

	if c.authToken == "" {
		if endpoint != loginEndpoint {
			return fmt.Errorf("client is not logged in")
		}
	} else {
		req.Header.Add("X-F5-Auth-Token", c.authToken)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if err = checkStatusCode(res); err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(model)
	if err != nil {
		return err
	}

	return nil
}

// Login attempts to retrieve an auth token from the API using the credentials the client was created with and returns nil.
// If login is unsuccessful an error is returned
func (c *F5Client) Login() error {
	loginArgs := map[string]string{
		"loginProviderName": "tmos",
		"username":          c.username,
		"password":          c.password,
	}
	loginBody, err := json.Marshal(loginArgs)
	if err != nil {
		return err
	}

	var loginResponse tokenResponse
	err = c.DoRequest(http.MethodPost, loginEndpoint, string(loginBody), &loginResponse)
	if err != nil {
		return err
	}

	if loginResponse.Token.Token == nil {
		return fmt.Errorf("couldn't get auth token from response")
	}

	// successful request, extract token
	c.authToken = *loginResponse.Token.Token
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
