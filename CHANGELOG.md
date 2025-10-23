# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- [#14] Extracted state diff constants from operator

## [v2.0.0] - 2025-10-10

**Breaking Changes ahead**

### Added
- [#7] added validation rules for dogus and config

### Changed
- [#5] **Breaking** use yaml format for blueprint and blueprint mask instead of json strings
- [#5] **Breaking** use absent bool instead of targetState for dogus and components
- [#5] **Breaking** only reference sensitive config by secret
- [#5] update dependencies
- [#7] **Breaking** flatten config structure to be less hierarchical
- [#7] **Breaking** use conditions instead of status phases
- [#7] **Breaking** rename `dryRun`to `stopped` to reflect the new operating mode of the blueprint operator better

### Removed
- [#9] **Breaking** remove components from blueprint

## [v1.3.0] - 2025-06-04
### Added
- [#3] blueprints now support additionalMounts for configMaps and secrets in dogus

## [v1.0.0] - 2025-02-26
### Added
- Initial release
- Move blueprint-operator structs from k8s-blueprint-operator to this new library