# Golang Data Types

Go is a statically typed language, meaning each variable must be declared with a specific type. Here are the main data types in Go:

## Basic Data Types

### Integer Types
- **int**: Size depends on the platform (32 or 64 bits)
- **int8**: 8-bit integer (-128 to 127)
- **int16**: 16-bit integer (-32768 to 32767)
- **int32**: 32-bit integer (-2147483648 to 2147483647)
- **int64**: 64-bit integer (-9223372036854775808 to 9223372036854775807)
- **uint**: Unsigned, size depends on the platform (32 or 64 bits)
- **uint8**: 8-bit unsigned integer (0 to 255)
- **uint16**: 16-bit unsigned integer (0 to 65535)
- **uint32**: 32-bit unsigned integer (0 to 4294967295)
- **uint64**: 64-bit unsigned integer (0 to 18446744073709551615)

### Floating Point Types
- **float32**: 32-bit floating-point number
- **float64**: 64-bit floating-point number

### Complex Types
- **complex64**: Complex number with float32 real and imaginary parts
- **complex128**: Complex number with float64 real and imaginary parts

### Boolean Type
- **bool**: Represents true or false

### String Type
- **string**: A sequence of characters

## Composite Data Types

### Array
A fixed-length sequence of elements of a single type.

### Slice
A dynamically-sized, flexible view into the elements of an array.

### Struct
A collection of fields.

### Map
An unordered collection of key-value pairs.

### Channel
A conduit through which goroutines communicate.

## Example

```go
package main

import "fmt"

func main() {
    // Integer
    var i int = 42
    // Floating point
    var f float64 = 3.14
    // String
    var s string = "Hello, Go!"
    // Boolean
    var b bool = true

    // Array
    var arr [3]int = [3]int{1, 2, 3}
    // Slice
    var slc []int = []int{4, 5, 6}
    // Struct
    type Person struct {
        Name string
        Age  int
    }
    var p Person = Person{Name: "Alice", Age: 30}
    // Map
    var m map[string]int = map[string]int{"one": 1, "two": 2}
    // Channel
    ch := make(chan int)

    fmt.Println(i, f, s, b)
    fmt.Println(arr)
    fmt.Println(slc)
    fmt.Println(p)
    fmt.Println(m)

    // Using channel
    go func() {
        ch <- 7
    }()
    fmt.Println(<-ch)
}
```

This example shows how to declare and initialize different basic and composite data types in Go, making it easier for beginners to understand.
