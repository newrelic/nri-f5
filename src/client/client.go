package client

import (
	"encoding/json"
	"errors"
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
	HTTPClient       *http.Client
	Username         string
	Password         string
	AuthToken        string
	BaseURL          string
	RequestSemaphore chan struct{}
}

const (
	loginEndpoint  = "/mgmt/shared/authn/login"
	logoutEndpoint = "/mgmt/shared/authz/tokens/%s"
)

var errDeleteAuthToken = errors.New("couldn't delete auth token")

// NewClient takes in arguments and creates and returns a client that will talk to the F5 API, or an error if one cannot be created
func NewClient(args *arguments.ArgumentList) (*F5Client, error) {
	var options []nrHttp.ClientOption

	if args.CABundleDir != "" {
		options = append(options, nrHttp.WithCABundleDir(args.CABundleDir))
	}
	if args.CABundleFile != "" {
		options = append(options, nrHttp.WithCABundleFile(args.CABundleFile))
	}
	options = append(options, nrHttp.WithTimeout(time.Duration(args.Timeout)*time.Second))

	httpClient, err := nrHttp.New(options...)
	if err != nil {
		return nil, err
	}

	return &F5Client{
		HTTPClient:       httpClient,
		Username:         args.Username,
		Password:         args.Password,
		AuthToken:        "",
		BaseURL:          "https://" + args.Hostname + ":" + strconv.Itoa(args.Port),
		RequestSemaphore: make(chan struct{}, args.MaxConcurrentRequests),
	}, nil
}

// Request is a shortcut for making a GET request without a request body
func (c *F5Client) Request(endpoint string, model interface{}) error {
	return c.DoRequest(http.MethodGet, endpoint, "", model)
}

// DoRequest makes a request to the given endpoint using the given request body, storing the result in the model if possible.
// An error is returned if either step cannot be completed.
func (c *F5Client) DoRequest(method, endpoint, body string, model interface{}) error {
	c.RequestSemaphore <- struct{}{}
	defer func() { <-c.RequestSemaphore }()

	req, err := http.NewRequest(method, c.BaseURL+endpoint, strings.NewReader(body))
	if err != nil {
		return err
	}

	if c.AuthToken == "" {
		if endpoint != loginEndpoint {
			return fmt.Errorf("client is not logged in")
		}
	} else {
		req.Header.Add("X-F5-Auth-Token", c.AuthToken)
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
	loginArgs := map[string]string{
		"loginProviderName": "tmos",
		"username":          c.Username,
		"password":          c.Password,
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
	c.AuthToken = *loginResponse.Token.Token
	return nil
}

// LogOut attempts to delete the auth token used in the actual session in order to not reach the maximum of 100 active tokens
// per user F5 Big Ip has (By default the tokens retrieved with the logIn call expire after 20 minutes)
func (c *F5Client) LogOut() error {
	logoutCall := fmt.Sprintf(logoutEndpoint, c.AuthToken)

	var logoutResponse tokenStruct
	err := c.DoRequest(http.MethodDelete, logoutCall, "", &logoutResponse)
	if err != nil {
		return err
	}

	if logoutResponse.Token == nil {
		return errDeleteAuthToken
	}

	c.AuthToken = ""
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
