# Changelog

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

The structure and content of this file follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [1.18.0] - 2023-03-07
### Added
- Added support for root fragments in filters such as `$.data[?(@.id == $.key)]`.
- "exists" is now an alias for the "has" filter operation.
- Added length, count, match, and search functions.
- Added `Nothing` as a value for comparison to return values where nothing is found.
- Added support no parenthesis around a filter so `[?@.x == 3]` is now valid.
- `alt.String()` now converts `[]byte`.
### Fixed
- Fix order of union with when final elements are not an `[]any`.

## [1.17.5] - 2023-02-19
### Added
- Added alt.Filter, a variation on alt.Match.
- Added the OmitEmpty option to oj, sen, pretty, and alt packages.
- Added the -o option for omit nil and empty to the oj command.

## [1.17.4] - 2023-02-02
### Fixed
- Fixed (preserve) order of JSONPath wildcard reflect elements.

## [1.17.3] - 2023-01-27
### Fixed
- Fixed (preserve) order of JSONPath filtered elements.

## [1.17.2] - 2023-01-15
### Fixed
- Fixed big number parsing.

## [1.17.1] - 2023-01-09
### Fixed
- Fixed the descent fragment use in the Modify() functions of the jp package.

## [1.17.0] - 2023-01-05
### Added
- Modify() functions added to the jp package.
- Added the `has` operator to the jp package scripts.

## [1.16.0] - 2023-01-02
### Added
- Remove() functions added to the jp package.
- jp.Set() operations now allow a union as the last fragment in an expression.

## [1.15.0] - 2022-12-16
### Added
- Added `jp.Script.Inspect()` to be able to get the details of a script.
- The parser callback function now allows `func(any)` in addition to `func(any) bool`.

## [1.14.5] - 2022-10-12
### Fixed
- alt.Builder Pop fixed for nested objects.

## [1.14.4] - 2022-08-11
### Fixed
- Private members that match a JSON element no longer cause a panic.

## [1.14.3] - 2022-06-12
### Fixed
- Returned `[]byte` from oj.Marshal and pretty.Marshal now copy the
  internal buffer instead of just returing