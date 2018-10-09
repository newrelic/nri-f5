package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/newrelic/nri-f5/src/arguments"
)

func Test_CreateClient(t *testing.T) {
	args := arguments.ArgumentList{
		Username: "testUser",
		Password: "testPass",
		Hostname: "testHost",
		Port:     1945,
	}

	client, err := NewClient(&args)
	assert.NoError(t, err)
	assert.Equal(t, "https://testHost:1945", client.baseURL)
	assert.Equal(t, "testUser", client.username)
	assert.Equal(t, "testPass", client.password)
	assert.Equal(t, "", client.authToken)
}

func Test_Login(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		t.Logf("Received request for %s", req.URL)
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/shared/authn/login" {
			requestBody, _ := ioutil.ReadAll(req.Body)
			bodyJSON := map[string]string{}
			_ = json.Unmarshal(requestBody, &bodyJSON)

			assert.Equal(t, "testUser", bodyJSON["username"])
			assert.Equal(t, "testPass", bodyJSON["password"])

			res.Write([]byte("{\"token\":{\"token\":\"this-is-a-token\"}}"))
		} else {
			assert.Equal(t, "this-is-a-token", req.Header.Get("X-F5-Auth-Token"))
			res.Write([]byte("{\"ok\":true}"))
		}
	}))
	defer func() { testServer.Close() }()

	client := F5Client{
		baseURL:    testServer.URL,
		username:   "testUser",
		password:   "testPass",
		httpClient: http.DefaultClient,
	}

	err := client.Request("/some-endpoint", nil)
	assert.Error(t, err)

	err = client.Login()
	assert.NoError(t, err)

	assert.Equal(t, "this-is-a-token", client.authToken)

	var okResp struct {
		OK *bool `json:"ok"`
	}
	err = client.Request("/some-endpoint", &okResp)
	assert.NoError(t, err)
	assert.Equal(t, true, *okResp.OK)
}

func Test_BadStatusCode(t *testing.T) {
	testResponse := http.Response{
		StatusCode: 404,
	}

	err := checkStatusCode(&testResponse)
	assert.Error(t, err)
}
