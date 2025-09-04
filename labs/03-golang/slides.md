---
title       : Getting started with Go.
author      : Haris Volos
description : This is an introduction to Go Lang.
keywords    : Go
marp        : true
paginate    : true
theme       : jobs
--- 

<style>

  .img-overlay-wrap {
    position: relative;
    display: inline-block; /* <= shrinks container to image size */
    transition: transform 150ms ease-in-out;
  }

  .img-overlay-wrap img { /* <= optional, for responsiveness */
    display: block;
    max-width: 100%;
    height: auto;
  }

  .img-overlay-wrap svg {
    position: absolute;
    top: 0;
    left: 0;
  }

  </style>

  <style>
  img[alt~="center"] {
    display: block;
    margin: 0 auto;
  }

</style>

<style>   

   .cite-author {     
      text-align        : right; 
   }
   .cite-author:after {
      color             : orangered;
      font-size         : 125%;
      /* font-style        : italic; */
      font-weight       : bold;
      font-family       : Cambria, Cochin, Georgia, Times, 'Times New Roman', serif; 
      padding-right     : 130px;
   }
   .cite-author[data-text]:after {
      content           : " - "attr(data-text) " - ";      
   }

   .cite-author p {
      padding-bottom : 40px
   }

</style>

<!-- _class: titlepage -->s: titlepage -->

![bg w:300 right:33%](assets/images/gophers.png)

# Lab: Getting Started with Go

---

# The Go programming language
- Modern
- Compact, concise, general-purpose
- Imperative, statically type-checked, dynamically type-safe
- Garbage-collected
- Compiles to native code, statically linked
- Fast compilation, efficient execution
- Designed by programmers for programmers!

---

# Install Go

golang.org/doc/install

- Install from binary distributions or build from source
- 32- and 64-bit x86 and ARM processors
- Windows, Mac OS X, Linux, and FreeBSD
- Other platforms may be supported by gccgo

## or ...
---

# Use CloudLab profile with Go installed

- Start a new experiment on CloudLab:
  - profile: `multi-node-cluster`
  - number of nodes: 1

- Ssh into `node0`, for example:

  ```
  ssh alice@amd227.utah.cloudlab.us
  ```

- Test your Go installation

  ```
  $ go version
  ```

---

# Part 1: Your first program

---

# Your first program

Change to your home directory

```bash
$ cd
```

Create a hello directory for your first Go source code

```bash
$ mkdir hello
$ cd hello
```

---

# Your first program

Put this code into hello.go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界!")
}
```

In this code, you 

- Declare a main package 

- Import the popular fmt package 

- Implement a main function to print a message to the console

---

# Your first program

Run the program:

```bash
$ go run hello.go
```

Output:

```
Hello, 世界!
```

---

# The go tool

The go tool is the standard tool for building, testing, and installing Go programs

```bash
$ go run hello.go           # Compile and run hello.go
```

---

# Package

A Go program consists of packages

A *package* consists of one or more source files (.go files)

Each source file starts with a package clause followed by declarations

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

By convention, all files belonging to a single package are located in a single directory

---

# Call code in an external package

Change `hello.go` so it looks like the following:

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

In this code, you:

- Import the package `quote`
- Add a call to its Go function

---

# Dependencies 

- Dependencies are packages which are required for your projects to run

- An import declaration is used to express a dependency on another package:

  ```go
  import "rsc.io/quote"
  ```

- Here, the importing package depends on the package `quote`

- The *import path* ("rsc.io/quote") uniquely identifies a package; multiple packages may have the same name, but they all reside at different directories

- By convention, the package name matches the last element of the import path (here: "quote")

- Exported functionality of the quote package is available via the *package qualifier*: `quote.Go`

---

# Dependency tracking

After your code imports the package, enable dependency tracking and get the package’s code to compile with

To add your code to its own module:

```bash
$ go mod init example/hello
```

- `go mod init` creates a new module, initializing the `go.mod` file
- `go.mod` describes the module and tracks the modules that provide those packages
- `example/hello` specifies a module path that serves as the module’s name

---

# Module

- A collection of related Go packages 
- Stored in a file tree with a go.mod file at its root
- Unique module path
- The unit of code interchange and versioning
- Semantically versioned

---

# go.mod

```
module example/hello

go 1.18

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
```

- Defines the module’s module path 
  - module path: the import path prefix of the module
- Defines the module's dependency requirements
  - dependency requirements: other modules needed for a successful build 

---

# Summary - The go tool

The go tool is the standard tool for building, testing, and installing Go programs

```bash
$ go run hello.go           # Compile and run hello.go

$ go test                   # Run tests

$ go build                  # Build and format the files in the current directory
$ go fmt                  

$ go mod init example/hello # Create a new module, initializing the go.mod file that describes it
                            # go build, go test, and other package-building commands add new
                            # dependencies to go.mod as needed

$ go get rsc.io/quote       # Fetch and install `quote`

$ go mod tidy               # Remove unused dependencies

$ go list -m all            # Print current module's dependencies 

```

---

# Part 2: Your first package

---

# Your first package

Put this code into `hello/greetings/greetings.go`:

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

In this code, you:
- Declare a `greetings` package to collect related functions
- Implement a `Hello` function to return the greeting

---
# Function syntax

![w:600 center](assets/images/function-syntax.png)


- Upper case names are exported: Name vs. name

  - A function whose name starts with a capital letter can be called by a function in another  package

---

# Call your code from another package

Change ``hello/hello.go`` so it looks like the following:

```go
package main

import (
    "fmt"

    "example/hello/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

---

# Call your code from another package

Run the program:

```bash
$ go run .
Hi, Gladys. Welcome!
```

---

# Return and handle an error

Handling errors is an essential feature of solid code

---

# Return an error

Change your code to return an error 

---

# diff hello/greetings/greetings.go

```diff
 package greetings

 import (
+    "errors"
     "fmt"
 )

 // Hello returns a greeting for the named person.
 func Hello(name string) (string, error) {
+    // If no name was given, return an error with a message.
+    if name == "" {
+        return "", errors.New("empty name")
+    }
+
     // If a name was received, return a value that embeds the name
     // in a greeting message.
     message := fmt.Sprintf("Hi, %v. Welcome!", name)
-    return message
     return message, nil
}
```

---

# Handle an error

Change your code to handle an error

---

# diff hello/hello.go

```diff
package main

 import (
     "fmt"
+    "log"

     "example/hello/greetings"
 )

 func main() {
+    // Set properties of the predefined Logger, including
+    // the log entry prefix and a flag to disable printing
+    // the time, source file, and line number.
+    log.SetPrefix("greetings: ")
+    log.SetFlags(0)

     // Request a greeting message.
-    message := greetings.Hello("Gladys")
+    message, err := greetings.Hello("")
+    // If an error was returned, print it to the console and
+    // exit the program.
+    if err != nil {
+        log.Fatal(err)
+    }
+
+    // If no error was returned, print the returned message
+    // to the console.
     fmt.Println(message)
}
```

---

# Return and handle an error

Run the program:

```bash
$ go run .
greetings: empty name
exit status 1
```

Now that you're passing in an empty name, you get an error

---

# Add a test

Testing your code during development can expose bugs

Go's built-in support for unit testing makes it easier to test as you go

Ending a file's name with `_test.go` tells the `go test` command that this file contains test functions

Let's add a test for the `Hello` function

---

# hello/greetings/greetings_test.go

```go
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
```

---

# go test

```bash
$ go test
PASS
ok      example/hello/greetings   0.364s

$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example/hello/greetings   0.372s
```

---

# Break a test

Let's break the `greetings.Hello` function to view a failing test

```diff
 // Hello returns a greeting for the named person.
 func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
-   message := fmt.Sprintf("Hi, %v. Welcome!", name)
+   message := fmt.Sprintf("Hi. Welcome!")
    return message, nil
 }
```

---

# Break a test

```bash
$ go test
--- FAIL: TestHelloName (0.00s)
    greetings_test.go:15: Hello("Gladys") = "Hi. Welcome!", <nil>, want match for `\bGladys\b`, nil
FAIL
exit status 1
FAIL    example/hello/greetings   0.182s
```

---

# Compile and install the application

`go run` compiles and runs a program; it doesn't generate a binary executable.

`go build` compiles the packages, along with their dependencies, but it doesn't install the results

`go install` compiles and installs the packages

---

# Exercises

### Return a random greeting

Change `Hello` so that instead of returning a single greeting every time, it returns one of several predefined greeting messages

### Return greetings for multiple people

Add a new function `Hellos` that returns greetings for multiple people

```go
// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error)
```
