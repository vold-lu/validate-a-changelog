# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.5.0] - 2025-10-28

### Changed

- docker: use alpine as base image, remove entrypoint.

## [0.4.0] - 2025-10-28

### Added

- Introduce cmd/lint-changelog.
- Parse changelog title.

### Changed

- Make parser more flexible regarding entry line.

### Fixed

- Handle malformed version with no section.

## [0.3.0] - 2025-10-28

### Added

- Introduce internal library.

### Changed

- cmd/validate-changelog: add new -allow-invalid-change-type-order flag.

## [0.2.0] - 2025-10-28

### Changed

- Validate that versions are in the right order.

## [0.1.2] - 2025-10-28

### Changed

- Handle case version entry is weirdly formatted.

## [0.1.1] - 2025-10-28

### Changed

- Make sure version is either "Unreleased" or valid SemVer.

## [0.1.0] - 2025-10-28

### Added

- Initial commit.