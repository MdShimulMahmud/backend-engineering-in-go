# Golang Variables

Variables in Go are used to store data that can be used and manipulated throughout a program. Here are the key concepts related to variables in Go:

## Declaring Variables

### Using `var` Keyword
You can declare a variable using the `var` keyword followed by the variable name and type.

```go
var name string
var age int
```

### Short Variable Declaration
You can also use the shorthand syntax `:=` to declare and initialize a variable in one line.

```go
name := "Alice"
age := 30
```

## Initializing Variables

Variables can be initialized at the time of declaration.

```go
var name string = "Alice"
var age int = 30
```

Using the shorthand syntax:

```go
name := "Alice"
age := 30
```

## Multiple Variable Declaration

You can declare and initialize multiple variables in a single line.

```go
var name, city string = "Alice", "Wonderland"
age, height := 30, 5.9
```

## Zero Values

If a variable is declared without an initial value, it is assigned a zero value by default.

- **int**: 0
- **float64**: 0.0
- **bool**: false
- **string**: ""

## Constants

Constants are declared using the `const` keyword and cannot be changed once set.

```go
const Pi = 3.14
const Greeting = "Hello, Go!"
```

## Example

```go
package main

import "fmt"

func main() {
    // Declaring variables
    var name string
    var age int

    // Initializing variables
    name = "Alice"
    age = 30

    // Short variable declaration
    city := "Wonderland"
    height := 5.9

    // Multiple variable declaration
    var country, language string = "Wonderland", "Go"
    isStudent, grade := true, 'A'

    // Constants
    const Pi = 3.14
    const Greeting = "Hello, Go!"

    fmt.Println(name, age, city, height, country, language, isStudent, grade, Pi, Greeting)
}
```

This example demonstrates how to declare, initialize, and use variables and constants in Go.
