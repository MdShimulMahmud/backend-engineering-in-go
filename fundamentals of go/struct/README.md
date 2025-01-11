# Structs in Go

## Introduction
In Go, a `struct` is a composite data type that groups together variables under a single name. These variables, known as fields, can have different types and are used to represent a record.

## Defining a Struct
To define a struct, use the `type` and `struct` keywords. Here is an example:

```go
type Person struct {
    Name string
    Age  int
    Address string
}
```

## Creating an Instance of a Struct
You can create an instance of a struct by using the struct name followed by curly braces:

```go
p := Person{
    Name: "John Doe",
    Age:  30,
    Address: "123 Main St",
}
```

## Accessing Struct Fields
You can access and modify the fields of a struct using the dot notation:

```go
fmt.Println(p.Name) // Output: John Doe
p.Age = 31
```

## Anonymous Structs
Go also supports anonymous structs, which are useful for one-time use:

```go
a := struct {
    Name string
    Age  int
}{
    Name: "Jane Doe",
    Age:  25,
}
```

## Struct Methods
You can define methods on structs to add behavior to them:

```go
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}
```

## Conclusion
Structs are a powerful feature in Go that allow you to create complex data types with multiple fields. They are essential for organizing and managing data in your Go programs.

## Further Reading
- [Go Structs](https://golang.org/doc/effective_go.html#structs)
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go Tour: Structs](https://tour.golang.org/moretypes/2)