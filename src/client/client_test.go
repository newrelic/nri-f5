package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/stretchr/testify/assert"
)

func Test_CreateClient(t *testing.T) {
	args := arguments.ArgumentList{
		Username: "testUser",
		Password: "testPass",
		Hostname: "testHost",
		Port:     1945,

		MaxConcurrentRequests: 1,
	}

	client, err := NewClient(&args)
	assert.NoError(t, err)
	assert.Equal(t, "https://testHost:1945", client.BaseURL)
	assert.Equal(t, "testUser", client.Username)
	assert.Equal(t, "testPass", client.Password)
	assert.Equal(t, "", client.AuthToken)
}

func Test_LogIn(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		t.Logf("Received request for %s", req.URL)
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/shared/authn/login" {
			requestBody, _ := ioutil.ReadAll(req.Body)
			bodyJSON := map[string]string{}
			_ = json.Unmarshal(requestBody, &bodyJSON)

			assert.Equal(t, "testUser", bodyJSON["username"])
			assert.Equal(t, "testPass", bodyJSON["password"])

			_, err := res.Write([]byte("{\"token\":{\"token\":\"this-is-a-token\"}}"))
			assert.NoError(t, err)
		} else {
			assert.Equal(t, "this-is-a-token", req.Header.Get("X-F5-Auth-Token"))
			_, err := res.Write([]byte("{\"ok\":true}"))
			assert.NoError(t, err)
		}
	}))
	defer testServer.Close()

	client := F5Client{
		BaseURL:          testServer.URL,
		Username:         "testUser",
		Password:         "testPass",
		HTTPClient:       http.DefaultClient,
		RequestSemaphore: make(chan struct{}, 1),
	}

	err := client.Request("/some-endpoint", nil)
	assert.Error(t, err)

	err = client.LogIn()
	assert.NoError(t, err)

	assert.Equal(t, "this-is-a-token", client.AuthToken)

	var okResp struct {
		OK *bool `json:"ok"`
	}
	err = client.Request("/some-endpoint", &okResp)
	assert.NoError(t, err)
	assert.Equal(t, true, *okResp.OK)
}

func Test_LogOut(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		t.Logf("Received request for %s", req.URL)

		if req.URL.String() == "/mgmt/shared/authz/tokens/a-valid-token" {
			res.WriteHeader(200)
			assert.Equal(t, "a-valid-token", req.Header.Get("X-F5-Auth-Token"))

			_, err := res.Write([]byte("{\"token\":\"a-valid-token\",\"another-param\":\"another\"}"))
			assert.NoError(t, err)
		} else {
			res.WriteHeader(401)
		}
	}))
	defer testServer.Close()

	client := F5Client{
		BaseURL:          testServer.URL,
		AuthToken:        "a-valid-token",
		HTTPClient:       http.DefaultClient,
		RequestSemaphore: make(chan struct{}, 1),
	}

	err := client.LogOut()
	assert.NoError(t, err)

	client.AuthToken = "an-invalid-token"
	err = client.LogOut()
	assert.Error(t, err)
}

func Test_BadStatusCode(t *testing.T) {
	testResponse := http.Response{
		StatusCode: 404,
	}

	err := checkStatusCode(&testResponse)
	assert.Error(t, err)
}
