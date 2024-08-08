
# Words Package

The `words` package provides utilities for converting strings between different
cases, handling initialisms, and splitting words based on specific rules.

The package aims to maximize performance and minimize dependencies. It doesn't use regular expressions.

## Installation

To install the package, run:

```sh
go get github.com/bartdeboer/words
```

## Functions

### ToSnakeCase

Converts a given string to snake_case.

```go
func ToSnakeCase(s string) string
```

Example:

```go
fmt.Println(ToSnakeCase("ThisIsATest")) // Output: this_is_a_test
```

### ToKebabCase

Converts a given string to kebab-case.

```go
func ToKebabCase(s string) string
```

Example:

```go
fmt.Println(ToKebabCase("ThisIsATest")) // Output: this-is-a-test
```

### Capitalize

Capitalizes the first letter of a given string and converts the rest to lowercase.

```go
func Capitalize(s string) string
```

Example:

```go
fmt.Println(Capitalize("example")) // Output: Example
```

### ToCapWords

Converts a given string to capitalized words, handling initialisms correctly.

```go
func ToCapWords(s string) string
```

Example:

```go
fmt.Println(ToCapWords("thisIsATest")) // Output: ThisIsATest
```

### ToMixedCase

Converts a given string to mixedCase, handling initialisms correctly.

```go
func ToMixedCase(s string) string
```

Example:

```go
fmt.Println(ToMixedCase("thisIsATest")) // Output: thisIsATest
```

### SplitWords

Splits a string into words based on transitions between alphanumeric and non-alphanumeric characters, as well as transitions between lowercase and uppercase characters.

```go
func SplitWords(s string) []string
```

Example:

```go
fmt.Println(SplitWords("ThisIsATest")) // Output: ["This", "Is", "A", "Test"]
```

## Tests

```sh
go test ./...
```

## Benchmarks

```sh
go test -bench=.
```
