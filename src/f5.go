package main

import (
	"os"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname     string `default:"localhost" help:"The hostname or IP of the F5 BIG IP device to monitor."`
	Port         int    `default:"443" help:"The port of the iControl API to connect to."`
	Username     string `default:"" help:"The username to connect to the F5 API with."`
	Password     string `default:"" help:"The password to connect to the F5 API with."`
	Timeout      int    `default:"30" help:"The number of seconds to wait before a request times out."`
	UseSSL       bool   `default:"true" help:"Whether or not to use SSL to connect to the API. The F5 API only allows connections using SSL."`
	CABundleFile string `default:"" help:"Alternative Certificate Authority bundle file"`
	CABundleDir  string `default:"" help:"Alternative Certificate Authority bundle directory"`
}

const (
	integrationName    = "com.newrelic.f5"
	integrationVersion = "0.1.0"
)

var (
	args argumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	exitOnErr(err)

	_, err = NewClient(&args)
	exitOnErr(err)

	exitOnErr(i.Publish())
}

func exitOnErr(err error) {
	if err != nil {
		log.Error("Encountered fatal error: %v", err)
		os.Exit(1)
	}
}
