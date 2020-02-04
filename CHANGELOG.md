# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

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

## v1.0.1

### Changed

- Remove CPU limits.

## v1.0.0

### Added

- `kiam` upstream helm chart `v3.4`

[v1.0.4]: https://github.com/giantswarm/kiam-app/pull/12
[v1.0.3]: https://github.com/giantswarm/kiam-app/pull/11
[v1.0.2]: https://github.com/giantswarm/kiam-app/pull/8
