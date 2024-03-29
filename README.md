[![GoDoc](https://pkg.go.dev/badge/github.com/blizzy78/consistent)](https://pkg.go.dev/github.com/blizzy78/consistent)


consistent
==========

A Go analyzer that checks that common constructs are used consistently.

**Example output**

```
test.go:9:10: declare the type of function parameters explicitly (consistent)
        _ = func(a, b int) {}
                ^
test.go:11:13: declare the type of function return values explicitly (consistent)
        _ = func() (a, b int) { return 1, 2 }
                   ^
test.go:13:6: use zero-value literal instead of calling new (consistent)
        _ = new(strings.Builder)
            ^
test.go:15:6: use slice literal instead of calling make (consistent)
        _ = make([]int, 0)
            ^
test.go:17:6: use lowercase digits in hex literal (consistent)
        _ = 0xABCDE
            ^
test.go:20:6: write common term in range expression on the left (consistent)
        _ = 1 < x && x < 10
            ^
test.go:22:6: use AND-NOT operator instead of AND operator with complement expression (consistent)
        _ = 1 & ^2
            ^
test.go:24:6: add zero before decimal point in floating-point literal (consistent)
        _ = .5
            ^
test.go:26:6: check if len is (not) 0 instead (consistent)
        _ = len([]int{}) > 0
            ^
test.go:29:7: separate cases with comma instead of using logical OR (consistent)
        case 1 < 2 || 3 < 4:
             ^
```

See below for a complete list of checks.

This analyzer works similar to [go-consistent], but with explicit configuration of checks instead of
auto-detection, which should make it faster.


List of Checks
--------------

**-params - check function/method parameter types**

```go
// -params explicit
func f(a int, b int) {
}

// -params compact
func f (a, b int) {
}
```


**-returns - check function/method return value types**

```go
// -returns explicit
func f() (a int, b int) {
	return 1, 2
}

// -returns compact
func f() (a, b int) {
	return 1, 2
}
```


**-typeParams - check type parameter types**

```go
// -typeParams explicit
func f[K any, V any]() {
}

// -typeParams compact
func f[K, V any]() {
}
```


**-funcTypeParams - check function type parameter types**

```go
// -funcTypeParams explicit
type f func(a int, b int)

// -funcTypeParams compact
type f func(a, b int)

// -funcTypeParams unnamed
type f func(int, int)
```


**-singleImports - check single import declarations**

```go
// -singleImports bare
import "fmt"

// -singleImports parens
import (
	"fmt"
)
```


**-newAllocs - check allocations using new**

```go
// -newAllocs literal
b := &strings.Builder{}

// -newAllocs new
b := new(strings.Builder)
```


**-makeAllocs - check allocations using make**

```go
// -makeAllocs literal
i := []int{}
m := map[int]int{}

// -makeAllocs make
i := make([]int, 0)
m := make(map[int]int)
```


**-hexLits - check upper/lowercase in hex literals**

```go
// -hexLits lower
h := 0xabcde

// -hexLits upper
h := 0xABCDE
```


**-rangeChecks - check range checks**

```go
// -rangeChecks left
if x > low && x < high {
}

// -rangeChecks center
if low < x && x < high {
}
```


**-andNOTs - check AND-NOT expressions**

```go
// -andNOTs and-not
x := a &^ b

// -andNOTs and-comp
x := a & ^b
```


**-floatLits - check floating-point literals**

```go
// -floatLits explicit
x := 0.5
y := -0.2

// -floatLits implicit
x := .5
y := -.2
```


**-lenChecks - check len/cap checks**

```go
// -lenChecks equalZero
if len(x) == 0 || len(y) != 0 {
}

// -lenChecks compareZero
if len(x) <= 0 || len(y) > 0 {
}

// -lenChecks compareOne
if len(x) < 1 || len(y) >= 1 {
}
```


**-switchCases - check switch case clauses**

```go
// -switchCases comma
case x == 1, y == 2

// -switchCases or
case x == 1 || y == 2
```


**-switchDefaults - check switch default clauses**

```go
// -switchDefaults last
switch x {
case "foo":
case "bar":
default:
}

// -switchDefaults first
switch x {
default:
case "foo":
case "bar":
}
```


**-emptyIfaces - check empty interfaces**

```go
// -emptyIfaces any
var x any

// -emptyIfaces iface
var x interface{}
```


**-slogAttrs - check log/slog argument types**

```go
// -slogAttrs bare
slog.Info("test", "value", 123, "foo", "bar")

// -slogAttrs attr
slog.Info("test", slog.Int("value", 123), slog.String("foo", "bar"))

// -slogAttrs consistent (both variants are valid)
slog.Info("test", "value", 123, "foo", "bar")
slog.Info("test", slog.Int("value", 123), slog.String("foo", "bar"))
```


**-labelsRegexp - check labels against regexp**

```go
// -labelsRegexp "^[a-z][a-zA-Z0-9]*$"
foo: // okay
FOO: // error
```


Running
-------

There are multiple ways to run the analyzer:

- Using Go directly:

  ```
  go run github.com/blizzy78/consistent/cmd/consistent@latest ARGS
  ```

- Install using Go, then running from $PATH:

  ```
  go install github.com/blizzy78/consistent/cmd/consistent@latest
  consistent ARGS
  ```

- Using Go vet:

  ```
  go install github.com/blizzy78/consistent/cmd/consistent@latest
  go vet -vettool=$(which consistent) ARGS
  ```


**Usage**

```
consistent: checks that common constructs are used consistently

Usage: consistent [-flag] [package]


Flags:
  -V	print version and exit
  -all
    	no effect (deprecated)
  -andNOTs value
    	check AND-NOT expressions (ignore/andNot/andComp) (default andNot)
  -c int
    	display offending line with this many lines of context (default -1)
  -cpuprofile string
    	write CPU profile to this file
  -debug string
    	debug flags, any subset of "fpstv"
  -emptyIfaces value
    	check empty interfaces (ignore/any/iface) (default any)
  -fix
    	apply all suggested fixes
  -flags
    	print analyzer flags in JSON
  -floatLits value
    	check floating-point literals (ignore/explicit/implicit) (default explicit)
  -funcTypeParams value
    	check function type parameter types (ignore/explicit/compact/unnamed) (default explicit)
  -hexLits value
    	check upper/lowercase in hex literals (ignore/lower/upper) (default lower)
  -json
    	emit JSON output
  -labelsRegexp value
    	check labels against regexp ("" to ignore) (default ^[a-z][a-zA-Z0-9]*$)
  -lenChecks value
    	check len/cap checks (ignore/equalZero/compareZero/compareOne) (default equalZero)
  -makeAllocs value
    	check allocations using make (ignore/literal/make) (default literal)
  -memprofile string
    	write memory profile to this file
  -newAllocs value
    	check allocations using new (ignore/literal/new) (default literal)
  -params value
    	check function/method parameter types (ignore/explicit/compact) (default explicit)
  -rangeChecks value
    	check range checks (ignore/left/center) (default left)
  -returns value
    	check function/method return value types (ignore/explicit/compact) (default explicit)
  -singleImports value
    	check single import declarations (ignore/bare/parens) (default bare)
  -slogAttrs value
    	check log/slog argument types (ignore/bare/attr/consistent) (default attr)
  -source
    	no effect (deprecated)
  -switchCases value
    	check switch case clauses (ignore/comma/or) (default comma)
  -switchDefaults value
    	check switch default clauses (ignore/last/first) (default last)
  -tags string
    	no effect (deprecated)
  -test
    	indicates whether test files should be analyzed, too (default true)
  -trace string
    	write trace log to this file
  -typeParams value
    	check type parameter types (ignore/explicit/compact) (default explicit)
  -v	no effect (deprecated)
```


License
-------

This package is licensed under the MIT license.



[go-consistent]: https://github.com/Quasilyte/go-consistent
