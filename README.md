
# godi

Simple dependency injection for Go.

# How to Use

## Register Singleton

```
package main

import (
    godi "github.com/bryan-t/godi"
    "io"
)

func main() {
    var buf bytes.Buffer
    godi.RegisterSingleton[io.Writer](&buf)
}

```

## Register Provider
Register a function that creates the service.

```
package main

import (
    godi "github.com/bryan-t/godi"
    "io"
)

func main() {
    godi.RegisterProvider[io.Writer](func()(io.Writer, error){
        var buf bytes.Buffer
        return &buf, nil
    })
}

```

## Get Service

```
package main

import (
    godi "github.com/bryan-t/godi"
    "io"
)

func main() {
    var buf bytes.Buffer
    godi.RegisterSingleton[io.Writer](&buf)
    writer, _ := godi.GetService[io.Writer]()
}

```

