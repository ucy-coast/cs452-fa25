Diff `hello/hello.go`

```diff
 package main

 import (
     "fmt"
     "log"

     "example.com/greetings"
 )

 func main() {
     // Set properties of the predefined Logger, including
     // the log entry prefix and a flag to disable printing
     // the time, source file, and line number.
     log.SetPrefix("greetings: ")
     log.SetFlags(0)

     // Request a greeting message.
-    message, err := greetings.Hello("")
+    message, err := greetings.Hello("Gladys")
     // If an error was returned, print it to the console and
     // exit the program.
     if err != nil {
         log.Fatal(err)
     }

     // If no error was returned, print the returned message
     // to the console.
     fmt.Println(message)
 }
```

#### New `hello/hello.go`

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("Gladys")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```


#### Old `hello/hello.go`

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```