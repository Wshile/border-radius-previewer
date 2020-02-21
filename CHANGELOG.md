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
  internal buffer instead of just returing it.

## [1.14.2] - 2022-06-03
### Added
- Added SameType test tool.

## [1.14.1] - 2022-05-31
### Fixed
- Removed dependency on external packages.

## [1.14.0] - 2022-04-08
### Added
- Added the JSONPath filter operation `in`.
- Added the JSONPath filter operation `empty`.
- Added the JSONPath filter operation `=~` for regex.

## [1.13.1] - 2022-03-19
### Fixed
- Fixed a case where a un-terminated JSON did not return an error.

## [1.13.0] - 2022-03-05
### Added
- Added jp.Expr.Has() function.
- Added jp.Walk to walk data and provide a the path and value for each
  element.

## [1.12.14] - 2022-02-28
### Fixed
- `[]byte` are encoded according to the ojg.Options.

## [1.12.13] - 2022-02-23
### Fixed
- For JSONPath (jp) reflection Get returns `has` value correctly for zero field values.

## [1.12.12] - 2021-12-27
### Fixed
- JSONPath scripts (jp.Script or [?(@.foo == 123)]) is now thread safe.

## [1.12.11] - 2021-12-10
### Fixed
- Parser reuse was no resetting callback and channels. It does now.

## [1.12.10] - 2021-12-07
### Added
- Added a delete option to the oj application.

## [1.12.9] - 2021-10-31
### Fixed
- Stuttering extracted elements when using the `-x` options has been fixed.

## [1.12.8] - 2021-09-21
### Fixed
- Correct unicode character is now included in error messages.

## [1.12.7] - 2021-09-14
### Fixed
- Typo in maxEnd for 32 bit architecture fixed.
- json.Unmarshaler fields in a struct correctly unmarshal.

## [1.12.6] - 2021-09-12
### Fixed
- Due to limitation (a bug most likely) in the stardard math package
  math.MaxInt64 can not be used on 32 bit architectures. Changes were
  made to work around this limitation.

- Embedded (Anonymous) pointers types now encode correctly.

### Added
- Support for json.Unmarshaler interface added.

## [1.12.5] - 2021-08-17
# Changed
- Updated to use go 1.17.

## [1.12.4] - 2021-08-06
### Fixed
- Setting an element in an array that does not exist now creates the array is the Nth value is not negative.

## [1.12.3] - 2021-08-01
### Fixed
- Error message on failed recompose was fixed to display the correct error message.
- Marshal of a non-pointer that contains a json.Marshaller that is not a poi