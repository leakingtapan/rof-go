[![Build Status](https://travis-ci.org/leakingtapan/rof-go.svg?branch=master)](https://travis-ci.org/leakingtapan/rof-go)
# Object Factory Go
Object Factory Go populates Go struct with randome value. This greatly saves time to create randomized object in unit testing. It carries the same philosophy as [Reflection Object Factory](https://github.com/leakingtapan/rof) in Java.

## Usage 
### Create Random Primitives
Random Go primitives can be created with simpily two lines of code:

```go
    package main
    
    import rof "github.com/leakingtapan/rof-go"
    
    func main() {
        // create random int
        // pointer is required for the varible
        var value int
        rof.Create(&value)
        
        // create random string
        // pointer is required for the variable
        var str string
        rof.Create(&str)
    }

```

### Create Random Composite Object
Fields of composite object will be populated recursively.

```go
    package main
    
    import rof "github.com/leakingtapan/rof-go"
    
    type Singer struct {
        FirstName string
        LastName string
        Songs []struct {
           Title string 
        } 
    }
    
    func main() {
        var s Singer
        rof.Create(&s)
    }

```

### Supported Built-in Types
* Primitives Types:
  - bool
  - int, int8, int16, int32, int64
  - uint, uint8, uint16, uint32, uint64
  - float32, float64
  - complex64, complex128
  - string
* Composite Types:
  - [time.Time](https://golang.org/pkg/time/#Time)

### Customize function providers
Each built-in types are creates using pre-built function similar to following signature:

```go
    func intFunc() int {
        return 128
    }

```

A custom function can be passed in using `rof.SetFunc` API:

```go
    rof.SetFunc(func() string {
        return "my own string"
    })
    
    var str string
    rof.Create(&str)
    fmt.Println(str) // "my own string"

```

Then this function will be used in following variable creation.
