<div style="text-align: center;">
<img src="./readmeassets/arguduck_300.png">
</div>

# ArguDuck
A quacking simple argmuent parser for Go.

## Installation
```bash
go get github.com/duck-fat-studios/ArguDcuk
```

## Argument Types

### Flag
 ```go
 Flag(name string, short string, help string, group ...string) (error, ArguDuckErrorString)
 ```

### Float
```go
Float(name string, short string, defaultValue float64, help string, group ...string) (error, ArguDuckErrorString)
```


### Int
```go
Int(name string, short string, defaultValue int, help string, group ...string) (error, ArguDuckErrorString) 
```

### String
```go
String(name string, short string, defaultValue string, help string, group ...string) (error, ArguDuckErrorString)
```


## Basic Usage
```go
package main

import (
    arguduck "github.com/duck-fat-studios/ArguDuck"
)

var (
    args = *arguduck.InitArguDuck()
)

func main() {
    setupArguments()

    fmt.Println(arguduck.Args["foo"])
    // >> Bar

}

func setupArguments() {
    args.SetAbout("this is the message")
    args.String("foo", "f", "bar", "Prints Bar", "FooBar")
    args.Parse()
}

```
## Error Handling
```go
err, errString := args.String("foo", "f", "bar", "FOOBAR!", "FooBar")

if err != nil {
    fmt.Printf("%s: %s\n", errString, err)
}
```