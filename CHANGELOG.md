# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## Changed

- Upgrade `kiam` version to 4.0.

## [1.6.0] - 2021-01-07

## [1.5.0] - 2020-09-15

### Added

- Setting gRPC environment variables

## [1.4.2] - 2020-08-27

### Changed

- Use deep liveness probe for kiam agent.
- Align charts with upstream.

## [1.4.1] - 2020-08-19

### Added

- Added monitoring and common labels.

## [1.4.0] - 2020-08-05

### Change

- Updated cert-manager API groups.
- Use latest cert-manager in integration tests.

## [1.3.1] - 2020-07-23

### Added

- Added worflows for automatic releases.

## Changed

- Upgrade `kiam` version to 3.6.

## [1.3.0] - 2020-07-06

### Added

- Set kiam region flag for STS endpoint

### Changed

- Upgrade architect-orb to 0.10.0

## [1.2.2] 2020-04-17

### Added

- Add configurable agent whitelist of proxy routes.

## [1.2.1] 2020-03-14

### Changed

- Interface name as string

## [1.2.0] 2020-03-11

### Changed

- Change name of the interfaces to manage for AWS CNI.


## [1.1.1] 2020-02-21

### Changed

- Use same registry configuration for parent and subchart.

## [1.1.0] 2020-02-04

### Changed

- Updated kiam app version to v3.5.

## [1.0.4] 2020-01-27

### Changed

- Fixed hook deletion policy for sub chart that labels kube-system namespace.

## [1.0.3] 2020-01-17

### Changed

- Updated server network policy labels to match server daemonset labels.
- Fixed hooks for sub chart that labels kube-system namespace.  

## [1.0.2] 2020-01-04

### Changed

- Updated manifests for Kubernetes 1.16.

## [1.0.1] 2019-12-27

### Changed

- Remove CPU limits.

## [1.0.0] 2019-10-25

### Added

- `kiam` upstream helm chart `v3.4`

[Unreleased]: https://github.com/giantswarm/kiam-app/compare/v1.6.0...HEAD
[1.6.0]: https://github.com/giantswarm/kiam-app/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/giantswarm/kiam-app/compare/v1.4.2...v1.5.0
[1.4.2]: https://github.com/giantswarm/kiam-app/compare/v1.4.1...v1.4.2
[1.4.1]: https://github.com/giantswarm/kiam-app/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/giantswarm/kiam-app/compare/v1.3.1...v1.4.0
[1.3.1]: https://github.com/giantswarm/kiam-app/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/giantswarm/kiam-app/compare/v1.2.2...v1.3.0
[1.2.2]: https://github.com/giantswarm/kiam-app/compare/tag/v1.2.2
[1.2.1]: https://github.com/giantswarm/kiam-app/compare/tag/v1.2.1
[1.2.0]: https://github.com/giantswarm/kiam-app/compare/tag/v1.2.0
[1.1.1]: https://github.com/giantswarm/kiam-app/compare/tag/v1.1.1
[1.1.0]: https://github.com/giantswarm/kiam-app/compare/tag/v1.1.0
[1.0.4]: https://github.com/giantswarm/kiam-app/compare/tag/v1.0.4
[1.0.3]: https://github.com/giantswarm/kiam-app/compare/tag/v1.0.3
[1.0.2]: https://github.com/giantswarm/kiam-app/compare/tag/v1.0.2
[1.0.1]: https://github.com/giantswarm/kiam-app/compare/tag/v1.0.1
[1.0.0]: https://github.com/giantswarm/kiam-app/compare/tag/v1.0.0
