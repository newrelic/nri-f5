# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

Unreleased section should follow [Release Toolkit](https://github.com/newrelic/release-toolkit#render-markdown-and-update-markdown)

## Unreleased

## v2.8.4 - 2025-07-03

### ⛓️ Dependencies
- Updated golang version to v1.24.4

## v2.8.3 - 2025-02-20

### ⛓️ Dependencies
- Updated golang patch version to v1.23.6

## v2.8.2 - 2025-01-30

### ⛓️ Dependencies
- Updated golang patch version to v1.23.5

## v2.8.1 - 2024-12-05

### ⛓️ Dependencies
- Updated golang patch version to v1.23.4

## v2.8.0 - 2024-10-10

### dependency
- Upgrade go to 1.23.2

### 🚀 Enhancements
- Upgrade integrations SDK so the interval is variable and allows intervals up to 5 minutes

## v2.7.7 - 2024-09-12

### ⛓️ Dependencies
- Updated golang version to v1.23.1

## v2.7.6 - 2024-07-04

### ⛓️ Dependencies
- Updated golang version to v1.22.5

## v2.7.5 - 2024-05-09

### ⛓️ Dependencies
- Updated golang version to v1.22.3

## v2.7.4 - 2024-04-11

### ⛓️ Dependencies
- Updated golang version to v1.22.2

## v2.7.3 - 2024-02-22

### ⛓️ Dependencies
- Updated github.com/newrelic/infra-integrations-sdk to v3.8.2+incompatible

## v2.7.2 - 2023-10-26

### ⛓️ Dependencies
- Updated golang version to 1.21

## 2.7.1 (2023-08-22)
### Fixed
- Removes config validation that prevents to run the integration without a custom CA certificates

## 2.7.0 (2023-06-06)
### Changed
- Upgrade Go version to 1.20

## 2.6.0 (2022-08-01)
### Fixed
Bumped the goversion used to build the integration. Upgrading the goversion to 1.18,

In particular in the new golang version the CommonName is no longer taken into consideration while validating certificates.
> The deprecated, legacy behavior of treating the CommonName field on X.509 certificates as a host name when no Subject Alternative Names are present is now disabled by default.

This could be an issue for users still relying on legacy `commonName` and not on `Subject Alternative Name`. In that case they would see an error message like:
```
[ERR] Encountered fatal error: Post [...] x509:  certificate relies on legacy Common Name field, use SANs instead
```

To overcome this issue the user should update the certificate relying on `Subject Alternative Name`. 
While the certificate is not updated, certificate validation could be disabled setting `--tls_insecure_skip_verify` to true.

## 2.5.3 (2022-06-20)
### Changed
- Updated dependencies
- Added support for RHEL 9 and Ubuntu 22.04

## 2.5.2 (2022-03-07)
### Changed
- Delete auth token after usage
- Add fallback calling new endpoint for device inventory
- Update Infrastracture SDK to v3.7.2

## 2.4.1 (2021-10-18)
### Added
Added support for more distributions:
- Debian 11
- Ubuntu 20.10
- Ubuntu 21.04
- SUSE 12.15
- SUSE 15.1
- SUSE 15.2
- SUSE 15.3
- Oracle Linux 7
- Oracle Linux 8

## 2.4.0 (2021-08-30)
### Added

Moved default config.sample to [V4](https://docs.newrelic.com/docs/create-integrations/infrastructure-integrations-sdk/specifications/host-integrations-newer-configuration-format/), added a dependency for infra-agent version 1.20.0

Please notice that old [V3](https://docs.newrelic.com/docs/create-integrations/infrastructure-integrations-sdk/specifications/host-integrations-standard-configuration-format/) configuration format is deprecated, but still supported.

## 2.3.1 (2021-06-11)
### Changed
- ARM support

## 2.3.0 (2021-05-10)
### Changed
- Update Go to v1.14.
- Migrate to Go Modules
- Update Infrastracture SDK to v3.6.7.
- Update other dependecies.

## 2.2.1 (2021-03-24)
### Changed
- Added arm packages and binaries 

## 2.2.0 (2020-09-02)
### Added
- `max_concurrent_connections` argument

## 2.1.1 (2019-11-18)
### Fixed
- Respect --metrics and --inventory flags

## 2.1.0 (2019-11-18)
### Changed
- Renamed the integration executable from nr-f5 to nri-f5 in order to be consistent with the package naming. **Important Note:** if you have any security module rules (eg. SELinux), alerts or automation that depends on the name of this binary, these will have to be updated.
## 2.0.2 - 2019-10-22
- Windows installer packaging

## 2.0.1 - 2019-07-23
- Removed unneeded JMX dependency

## 2.0.0 - 2019-05-06
### Changed
- Updated SDK
- Added more unique IDAttributes to entities

## 1.0.3 - 2019-03-19
### Fixed
- Rename metric pool.currentConnections to pool.sessions

## 1.0.2 - 2019-02-26
### Fixed
- Fix definition file to correct inventory prefix

## 1.0.1 - 2019-02-05
### Changed
- Bumped the version to 1.0.1

## 0.1.2 - 2019-02-04
### Fixed
- Changed the protocol version

## 0.1.1 - 2018-11-18
### Added
- Added metadata with URL of collection

## 0.1.0 - 2018-11-05
### Added
- Initial version: Includes Metrics and Inventory data
