# New Relic Infrastructure Integration for F5 BIG-IP 

Reports status and metrics for F5 BIG-IP

See our [documentation web site](https://docs.newrelic.com/docs/integrations/host-integrations/host-integrations-list/f5-monitoring-integration) for more details.

## Requirements

None

## Installation

* Download an archive file for the `f5` Integration
* Extract `f5-definition.yml` and the `bin` directory into `/var/db/newrelic-infra/newrelic-integrations`
* Add execute permissions for the binary file `nri-f5` (if required)
* Extract `f5-config.yml.sample` into `/etc/newrelic-infra/integrations.d`

## Usage

To run the F5 BIG-IP integration, you must have the agent installed (see [agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

To use the integration, first rename `f5-config.yml.sample` to `f5-config.yml`, then configure the integration
by editing the fields in the file. 

You can view your data in Insights by creating your own NRQL queries. To do so, use the **F5BigIpNodeSample**, **F5BigIpPoolMemberSample**, **F5BigIpPoolSample**, **F5BigIpSystemSample**, **F5BigIpVirtualServerSample** events in Insights.

## Compatibility

* Supported OS: No limitations
* F5 BIG-IP 11.6+

## Integration Development usage

Assuming you have the source code, you can build and run the integration locally

* Go to the directory of the F5 Integration and build it
```
$ make
```

* The command above will execute tests for the F5 BIG-IP integration and build an executable file called `nri-f5` in the `bin` directory
```
$ ./bin/nri-f5 --help
```

For managing external dependencies, the [govendor tool](https://github.com/kardianos/govendor) is used. It is required to lock all external dependencies to a specific version (if possible) in the vendor directory.
