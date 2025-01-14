# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.12] - 2019-10-22

### Changed
- Names for logical branches are inherited from reference names where available.

### Added
- Option to `$inheritSchemaFromExamples` when type and properties are missing.

### Fixed
- Invalid property name when stripping parent prefix.

## [0.4.11] - 2019-10-15

### Fixed
- Dereference of `nil` `*json.RawMessage` in additional and pattern properties.

## [0.4.10] - 2019-10-15

### Added
- Support for `x-omitempty` vendor extension.
- Vendor extensions documentation in README.md.

## [0.4.9] - 2019-10-10

### Added
- Option to disable null `additionalProperties` rendering.

## [0.4.8] - 2019-10-09

### Changed
- Constant/enum type names are prefixed with parent property if available.
- Constant/enum item names are prefixed with parent type instead of path. 
- Schema title is used for struct name instead of path where available.

### Added
- Optional support for `x-nullable` (Swagger 2.0), `nullable` (OpenAPI 3.0) properties to enable nullability.

### Fixed
- Type marker stripping from path was affecting regexps too.

## [0.4.7] - 2019-10-05

### Added
- Support for `x-go-type` as object (`go-swagger` format).
- Builder option `$ignoreXGoType` to disregard `x-go-type` hints.
- Builder option `$withZeroValues` to use pointer types to avoid zero-value ambiguity.
- Builder option `$ignoreNullable` to enable `omitempty` for nullable properties.
- Automated API docs.
- Change log.

### Fixed
- Missing processing for `null` `additionalProperties`.
- Processing for nullable (`[null, <other>]`) types.

### Changed
- Multi-type schemas are decomposed into multiple structures.
- Schema title is used for structure name if available.

## [0.4.6] - 2019-10-01

### Added
- More tests.
- CI upgrade.

### Fixed
- Code-style issues.

## [0.4.5] - 2019-07-11

### Added
- Type inference from enum values.

### Changed
- Trivial nesting removed. 

## [0.4.4] - 2019-07-08

### Fixed
- Removed unnecessary regexp dependency, #7.

[0.4.12]: https://github.com/swaggest/go-code-builder/compare/v0.4.11...v0.4.12
[0.4.11]: https://github.com/swaggest/go-code-builder/compare/v0.4.10...v0.4.11
[0.4.10]: https://github.com/swaggest/go-code-builder/compare/v0.4.9...v0.4.10
[0.4.9]: https://github.com/swaggest/go-code-builder/compare/v0.4.8...v0.4.9
[0.4.8]: https://github.com/swaggest/go-code-builder/compare/v0.4.7...v0.4.8
[0.4.7]: https://github.com/swaggest/go-code-builder/compare/v0.4.6...v0.4.7
[0.4.6]: https://github.com/swaggest/go-code-builder/compare/v0.4.5...v0.4.6
[0.4.5]: https://github.com/swaggest/go-code-builder/compare/v0.4.4...v0.4.5
[0.4.4]: https://github.com/swaggest/go-code-builder/compare/v0.4.3...v0.4.4
