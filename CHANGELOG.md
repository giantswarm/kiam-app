# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Set kiam region flag for STS endpoint

### Changed

- Upgrade architect-orb to 0.10.0

## [v1.2.2] 2020-04-17

### Added

- Add configurable agent whitelist of proxy routes.

## [v1.2.1] 2020-03-14

### Changed

- Interface name as string

## [v1.2.0] 2020-03-11

### Changed

- Change name of the interfaces to manage for AWS CNI.


## [v1.1.1] 2020-02-21

### Changed

- Use same registry configuration for parent and subchart.

## [v1.1.0] 2020-02-04

### Changed

- Updated kiam app version to v3.5.

## [v1.0.4] 2020-01-27

### Changed

- Fixed hook deletion policy for sub chart that labels kube-system namespace.

## [v1.0.3] 2020-01-17

### Changed

- Updated server network policy labels to match server daemonset labels.
- Fixed hooks for sub chart that labels kube-system namespace.  

## [v1.0.2] 2020-01-04

### Changed

- Updated manifests for Kubernetes 1.16.

## [v1.0.1] 2019-12-27

### Changed

- Remove CPU limits.

## [v1.0.0] 2019-10-25

### Added

- `kiam` upstream helm chart `v3.4`

[v1.2.2]: https://github.com/giantswarm/kiam-app/releases/tag/v1.2.2
[v1.2.1]: https://github.com/giantswarm/kiam-app/releases/tag/v1.2.1
[v1.2.0]: https://github.com/giantswarm/kiam-app/releases/tag/v1.2.0
[v1.1.1]: https://github.com/giantswarm/kiam-app/releases/tag/v1.1.1
[v1.1.0]: https://github.com/giantswarm/kiam-app/releases/tag/v1.1.0
[v1.0.4]: https://github.com/giantswarm/kiam-app/releases/tag/v1.0.4
[v1.0.3]: https://github.com/giantswarm/kiam-app/releases/tag/v1.0.3
[v1.0.2]: https://github.com/giantswarm/kiam-app/releases/tag/v1.0.2
[v1.0.1]: https://github.com/giantswarm/kiam-app/releases/tag/v1.0.1
[v1.0.0]: https://github.com/giantswarm/kiam-app/releases/tag/v1.0.0
