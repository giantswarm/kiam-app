# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Allow running without iptables mode
- Add cilium local redirect policy
- Move to Cilium Network Policies
- Move to main catalog

## [2.6.0] - 2022-11-18

- Support for running behind a proxy.
  - `HTTP_PROXY`,`HTTPS_PROXY` and `NO_PROXY` are set as environment variables in deployment if defined in `values.yaml`.
- Support for using `cluster-apps-operator` specific `cluster.proxy` values.

## [2.5.1] - 2022-11-15

### Fixed

- Allow `whiteListRouteRegexp` to default to `/latest/*`.

## [2.5.0] - 2022-08-10

## [2.4.0] - 2022-08-09

### Changed

- Change the kube-system annotation job's Restart Policy to `OnFailure`.

## [2.3.0] - 2022-03-21

### Added

- Add VerticalPodAutoscaler CR.

## [2.2.0] - 2022-03-08

### Changed

- Updated `whiteListRouteRegexp` to default to `/latest/meta-data/placement/availability-zone`

## [2.1.5] - 2022-02-28

## [2.1.4] - 2022-02-25

## [2.1.3] - 2022-02-25

## [2.1.2] - 2022-02-25

### Fixed

- Merged two release workflows into one to handle both tags

## [2.1.1] - 2022-02-25

### Added

- Build script to generate an IRSA compatible version of each release

## [2.1.0] - 2022-01-03

### Changed

- Upgrade `kiam` version to 4.2.
- Increase AWS session duration to `60m`.
- Increase AWS session refresh to `10m`.

## [2.0.0]

### Changed

- Upgrade `kiam` version to 4.1.
- Update RBAC API version from `v1beta1` to `v1`.
- Add `kind: Issuer` and `group: cert-manager.io` to `Certificate` templates.

## [1.7.1] - 2021-03-26

### Changed

- Set docker.io as the default registry

## [1.7.0] - 2021-01-21

- Add taint tolerations for kiam agent and kiam server.

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

[Unreleased]: https://github.com/giantswarm/kiam-app/compare/v2.6.0...HEAD
[2.6.0]: https://github.com/giantswarm/kiam-app/compare/v2.5.1...v2.6.0
[2.5.1]: https://github.com/giantswarm/kiam-app/compare/v2.5.0...v2.5.1
[2.5.0]: https://github.com/giantswarm/kiam-app/compare/v2.4.0...v2.5.0
[2.4.0]: https://github.com/giantswarm/kiam-app/compare/v2.3.0...v2.4.0
[2.3.0]: https://github.com/giantswarm/kiam-app/compare/v2.2.0...v2.3.0
[2.2.0]: https://github.com/giantswarm/kiam-app/compare/v2.1.5...v2.2.0
[2.1.5]: https://github.com/giantswarm/kiam-app/compare/v2.1.4...v2.1.5
[2.1.4]: https://github.com/giantswarm/kiam-app/compare/v2.1.3...v2.1.4
[2.1.3]: https://github.com/giantswarm/kiam-app/compare/v2.1.2...v2.1.3
[2.1.2]: https://github.com/giantswarm/kiam-app/compare/v2.1.1...v2.1.2
[2.1.1]: https://github.com/giantswarm/kiam-app/compare/v2.1.0...v2.1.1
[2.1.0]: https://github.com/giantswarm/kiam-app/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/giantswarm/kiam-app/compare/v1.7.1...v2.0.0
[1.7.1]: https://github.com/giantswarm/kiam-app/compare/v1.7.0...v1.7.1
[1.7.0]: https://github.com/giantswarm/kiam-app/compare/v2.0.0...v1.7.0
[2.0.0]: https://github.com/giantswarm/kiam-app/compare/v1.6.0...v2.0.0
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
