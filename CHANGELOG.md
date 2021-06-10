# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

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
