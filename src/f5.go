package main

import (
	"os"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
)

const (
	integrationName    = "com.newrelic.f5"
	integrationVersion = "0.1.0"
)

var (
	args arguments.ArgumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	exitOnErr(err)

	client, err = client.NewClient(&args)
	exitOnErr(err)

	// set up and run goroutines for each entity

	// go 

	exitOnErr(i.Publish())
}

func exitOnErr(err error) {
	if err != nil {
		log.Error("Encountered fatal error: %v", err)
		os.Exit(1)
	}
}
