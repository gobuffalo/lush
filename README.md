# Lush

[![Build Status](https://dev.azure.com/markbates/buffalo/_apis/build/status/gobuffalo.lush?branchName=master)](https://dev.azure.com/markbates/buffalo/_build/latest?definitionId=50&branchName=master)[![GoDoc](https://godoc.org/github.com/gobuffalo/lush?status.svg)](https://godoc.org/github.com/gobuffalo/lush)

Lush is an embeddable scripting language for Go with minimal dependencies.

Lush will provide the basis for [https://github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) v4.

**WIP**: Lush is currently a work in progress. We would love your feedback (and PR's), but be warned that things are liable to shift.

---

# Table of Contents

* [CLI](#cli)
* [Comments](#comments)
* [Identifiers](#identifiers)
* [Variables](#variables)
* [Strings](#strings)
* [Numbers](#numbers)
* [Booleans](#booleans)
* [If Statements](#if)
* [Operators](#operators)
* [Arrays](#arrays)
* [Maps](#maps)
* [For Loops](#for)
* [Functions](#functions)
* [Returns](#returns)
* [Scope](#scope)
* [Custom Helpers](#custom-helpers)
* [Errors](#errors)
* [Goroutines](#goroutines)
* [Calling](#calling)
* [Helpers](#helpers)
* [Imports](#imports)
* [Built-ins](#builtins)

---

## [CLI](#cli)

### Installation

```bash
$ go get -u github.com/gobuffalo/lush/cmd/lush
```

### Usage

#### Run Lush Scripts

```bash
$ lush run ./path/to/file.lush
```

#### Format Lush Scripts

```bash
$ lush fmt ./path/to/file.lush
```

##### Flags

* `-w` - write result to (source) file instead of stdout
* `-d` - display diffs instead of rewriting files

#### Print AST

```bash
$ lush ast ./path/to/file.lush
```

---

## [Comments](#comments)

```lush
// my comment

// (Deprecated):
# my comment
```

---

## [Identifiers](#identifiers)

* must start with `[a-zA-z]`
* `[0-9]` is allowed after first char
* `_` is allowed after first char

```lush
// Valid
a := 1
a1 := 1
ab1 := 1
ab1a := 1
ab1c2 := 1
a_1 := 1
a_1_b := 1
a_B_1_C_2 := 1

// Invalid (not a complete list - just examples)
true := 1
false := 1
nil := 1
_x := 1
1x := 1
x! := 1
```

---

## [Variables](#variables)

**TODO:**

* Allow multiple variable assignment: `x, y := a, b`

### `:=` (Declaration and Assignment)

The `:=` operator will create a new variable and assign it the given value. If the variable already exists in the context, an error will be returned

```lush
a := 1
```

### `let` (Declaration and/or Assignment)

The `let` keyword will create a new variable, if it doesn't already exist, and assign the given value to it. If the variable already exists its value will be replaced with the new value.

```lush
let a = 1
```

### `=` (Assignment)

The `=` will assign a new value to an existing variable. If the variable doesn't already exist in the context an error will be returned.

When assigning to an existing variable it's value will be modified in all parent contexts that contain that variable.

```lush
x := 0

func() {
  if true {
    x = 42
  }
}()

return x // 42
```

---

## [Strings](#strings)

* `"my string"` - Interpreted string literal
* <code>`multiline string`</code> - Multiline string literal

```lush
"foo"
`this
is
a
multi
line
string`
```

---

## [Numbers](#numbers)

* `42`, `-42` - `int` values
* `4.2`, `-4.2` - `float64` values

---

## [Booleans](#booleans)

```lush
true
false
```

---

## Nil

```lush
nil
```

---

## [`if` Statements](#if)

```lush
// with bool
if (true) {
  // do work
}

// check equality
if (a == b) {
  // do work
}

// optional parens
if a == b {
  // do work
}

// var declaration as pre-condition
if a := 1; a == 1 {
  return true
}
```

### `else` Statement

An `if` statement can only have one `else` statement attached to it.

```lush
if (false) {
  fmt.Println("in if")
} else {
  fmt.Println("in else")
}
```

### `else if` Statements

An `if` statement can have `N` `else if` statements.

```lush
if false {
  fmt.Println("in if")
} else if (1 == 2) {
  fmt.Println("in else")
} else if true {
  fmt.Println("2 == 2")
} else {
  fmt.Println("in other else")
}
```

---

## [Operators](#operators)

* `&&` - Requires both expressions to be `true`.
* `||` - Requires one of the expressions to be `true`.
* `==` - Equality operator. Uses [`github.com/google/go-cmp/cmp`](https://godoc.org/github.com/google/go-cmp/cmp) for comparison.
* `!=` - Inequality operator. Uses [`github.com/google/go-cmp/cmp`](https://godoc.org/github.com/google/go-cmp/cmp) for comparison.
* `+` - Adds statements together. Supports number types (`int`, `float64`), string concatenation, and array appending.
* `-` - Subtracts statements. Supports only number types (`int`, `float64`)
* `/` - Divides statements. Supports only number types (`int`, `float64`)
* `*` - Multiplys statements. Supports only number types (`int`, `float64`)
* `%` - Modulus operator. Supports only number types (`int`, `float64`)
* `>` - Greater than operator. Supports all types with a `string` comparison.
* `<` - Less than operator. Supports all types with a `string` comparison.
* `<=` - Less than or equal operator. Supports all types with a `string` comparison.
* `>=` - Greater than or equal operator. Supports all types with a `string` comparison.
* `~=` - Regular expression operator. Supports `string` type only. `"abc" ~= "^A"`

---

## [Arrays](#arrays)

Arrays are backed by `[]interface{}` so they can contain any known type.

### Defining

```lush
a := [1, "a", true, [4, 5, nil], {"x": "y", "z": "Z"}]
```

### Iterating

There are multiple ways to iterate over an array.

```lush
// `x` is `index` of array
for x := range [1, 2, 3] {
  // do work
}

// `i` is index of loop, `x` is `value` of array
for i, x := range myArray {
  // do work
}

// `x` is `value` of array
for (x) in [1, 2, 3] {
  fmt.Println(x, y)
}

// `i` is index of loop, `x` is `value` of array
for (i, x) in myArray {
  // do work
}
```

---

## [Maps](#maps)

Maps are backed by `map[interface{}]interface{}` and can contain any legal Go value.

### Defining

```lush
j := {"a": "b", "h": 1, "foo": "bar", "y": func(x) {}}
```

### Iterating

```lush
// `v` is `value` of the map
for v := range {foo: "bar", "x": 1} {
  // do work
}

// `k` is key, `v` is `value` of the map
for k, v := range myMap {
  // do work
}

// `v` is `value` of the map
for (v) in [1, 2, 3] {
  fmt.Println(v)
}

//  `k` is key, `v` is `value` of the map
for (k, v) in myMap {
  // do work
}
```

---

## [For Loops](#for)

### `break`

The `break` keyword works with both Maps and Arrays.

```lush
for x := range [1, 2, 3] {
  // do work
  break
}

for v := range {foo: "bar", "x": 1} {
  // do work
  break
}
```

### `continue`

The `continue` keyword works with both Maps and Arrays.

```lush
for x := range [1, 2, 3] {
  // do work
  break
}

for v := range {foo: "bar", "x": 1} {
  // do work
  break
}
```

### Infinite

```lush
for {
  if (i == nil) {
    let i = 0
  }

  i = (i + 1)

  if (i == 4) {
    return i // breaks the loop and returns `i`
  }
}
```

### Iterators

The `range` keyword supports an `Iterator` interface.

```go
type Iterator interface {
  Next() interface{}
}
```

---

## [Functions](#functions)

Functions can be defined using the `func` keyword. They can take and return `N` arguments.

```lush
myFunc := func(x) {
  return strings.ToUpper(x)
}

x := myFunc("hi")

fmt.fmt.Printlnln(x) // HI
```

---

## [Returns](#returns)

The `return` keyword can return `N` number of items.

```lush
return 1, "A", true
```

### Inside Functions

A `return` inside of a function can be used to return a value from the function. This will **not** stop the execution of the script.

```lush
f := func() {
  return 42 // returns 42 to when the function is executed
}
f() // does not exit the script
```

### Outside Functions

A `return` outside of a function will be returned automatically, and the execution of the script will stop.

```lush
if true {
  return 42 // returns 42 to the caller
}
```

---

## [Scope](#scope)

When a new code block, defined by `{ ... }`, is called, a new clone of the current Context is created for that block.

---

## [Custom Helpers](#custom-helpers)

Custom helper functions can be added to the Context before the script is executed.

```go
c := ast.NewContext(context.Background(), os.Stdout)
c.Set("myFunc", func(s string) string {
  return strings.ToUpper(s)
})
```

```lush
x := "a string"
fmt.fmt.Printlnln(myFunc(x)) // A STRING
```

### Optional Map

Custom helper functions can also take an optional map of type `map[string]interface{}` as the last, or second to last, argument.

```go
c := ast.NewContext(context.Background(), os.Stdout)
c.Set("myFunc", func(s string, opts map[string]interface{}) string {
  if _, ok := opts["lower"]; ok {
    return strings.ToLower(s)
  }
  return strings.ToUpper(s)
})
```

```lush
x := "A String"
fmt.fmt.Printlnln(myFunc(x)) // A STRING
fmt.fmt.Printlnln(myFunc(x, {lower: true})) // a string
```

### Optional Context

Custom helper functions can gain access to the current Context by accepting it as an optional last argument.

```go
c := ast.NewContext(context.Background(), os.Stdout)
c.Set("myFunc", func(s string, c *ast.Context) (string, error) {
  if c.Block != nil {
    res, err := c.Block.Exec(c)
    if err != nil {
      return "", err
    }
    return fmt.Sfmt.Println(res), nil
  }
  return strings.ToUpper(s), nil
})
```

```lush
x := "A String"
fmt.fmt.Printlnln(myFunc(x)) // A STRING

s := myFunc(x) {
  return "another string"
}
fmt.fmt.Printlnln(s) // another string
```

---

## [Errors](#errors)

This is set by default when creating a new `Context`. It is a function mapped to [`fmt#Errorf`](https://godoc.org/fmt#Errorf).

```lush
return error("stop %s", "dragging my heart around")
```

---

## [Goroutines](#goroutines)

When using Goroutines within a Lush script, Lush will wait until all Goroutines have completed before exiting.

```lush
go func() {
  let d = time.ParseDuration("1s")

  i := 0

  for {
    fmt.Println("xxx")

    time.Sleep(d)

    i = (i + 1)

    if (i == 5) {
      break
    }
  }
}()

```

---

## [Calling](#calling)

---

## [Helpers](#helpers)

---

## [Imports](#imports)

Imports differ from [helpers](#helpers) in that helpers are **automatically** available inside of a script, whereas imports need to be explicitly included.

For example to use the [built-in](#builtins) implementation of the `fmt` package, you would first `import "fmt"`.

```lush
import "fmt"

fmt.Println("foo")
```

See [`github.com/gobuffalo/lush/builtins#Available`](https://godoc.org/github.com/gobuffalo/lush/builtins#Available) for a full list of packages that are available for import.

### Adding Imports

To make something available for import, it must first be added to [`github.com/gobuffalo/lush/ast#Context.Imports`](https://godoc.org/github.com/gobuffalo/lush/ast#Context.Imports).

```go
c := ast.NewContext(context.Background(), os.Stdout)

c.Imports.Store("mypkg", mypkg{})
```

### CLI Imports

When running a Lush script using the CLI tool, the `-import` flag allows for making the [built-in](#builtins) package implementations available for importing into the script.

Of example the [`github.com/gobuffalo/lush/builtins#OS`](https://godoc.org/github.com/gobuffalo/lush/builtins#OS) built-in isn't include by default for security/safety reasons. To allow this to be imported by the script you can use the `-import` flag to allow access.

```bash
$ lush run -import os ./examples/big.lush
```

---

## [Built-ins](#builtins)

The [`github.com/gobuffalo/lush/builtins`](https://godoc.org/github.com/gobuffalo/lush/builtins) package provides implementations of a small set of the Go standard library.

