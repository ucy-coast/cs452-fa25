# Lab: Getting Started with Go

In this lab tutorial, you'll get a brief introduction to Go programming.

## Prerequisites

### Setting up the Experiment Environment in Cloudlab

For this tutorial, you will be using a CloudLab profile that comes with the latest version of Go. 

Start a new experiment on CloudLab using the `multi-node-cluster` profile in the `UCY-COAST-TEACH` project, configured with a single physical machine node. 

Open a remote SSH terminal session to `node0`.

Verify that the profile has a working installation of Go by typing the following command:

```
$ go version
```

Confirm that the command prints the installed version of Go. If you don't have Go installed then just follow the [download and install](https://go.dev/doc/install) steps.

## Part 1: Your first program

In this part, you will get started with a simple "Hello, World" program and learn a bit about Go code, tools, packages, and modules.

### Write some code

Get started with Hello, World.

1.  Open a remote SSH terminal session to `node0` and cd to your home directory

    ```
    $ cd
    ```

2.  Create a `$HOME/hello` directory for your first Go source code.

    For example, use the following commands:

    ```
    $ mkdir hello
    $ cd hello
    ```

3.  Enable dependency tracking for your code.

    When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That `go.mod` file stays with your code, including in your source code repository.

    To enable dependency tracking for your code by creating a `go.mod` file, run the `go mod init` command, giving it the name of the module your code will be in. The name is the module's module path.

    In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be `github.com/mymodule`. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. For more about naming a module with a module path, see [Managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module).

    For the purposes of this tutorial, just use `example/hello`.

    ```
    $ go mod init example/hello
    go: creating new go.mod: module example/hello
    ```

4.  In your text editor, create a file `hello.go` in which to write your code.

5.  Paste the following code into your `hello.go` file and save the file.

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, Go!")
    }
    ```

    This is your Go code. In this code, you:

    - Declare a main package. A package is a way to group functions, and it's made up of all the files in the same directory. The first statement in a Go source file must be package name. Executable commands must always use `package main`. 
    - Import the popular [`fmt` package](https://pkg.go.dev/fmt/), which contains functions for formatting text, including printing to the console. This package is one of the [standard library](https://pkg.go.dev/std) packages you got when you installed Go.
    - Implement a `main` function to print a message to the console. A `main` function executes by default when you run the `main` package.
  
6.  Run your code to see the greeting.

    ```
    $ go run .
    Hello, Go!
    ```

    The `go run` command is one of many `go` commands you'll use to get things done with Go. Use the following command to get a list of the others:

    ```
    $ go help
    ```

### Call code in an external package

When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.

1.  Make your printed message a little more interesting with a function from an external module.

    1. Visit pkg.go.dev and [search for a "quote" package](https://pkg.go.dev/search?q=quote).
    2. Locate and click the [`rsc.io/quote`](https://pkg.go.dev/rsc.io/quote) package in search results (if you see rsc.io/quote/v4, ignore it for now).
    3. In the **Documentation** section, under **Index**, note the list of functions you can call from your code. You'll use the Go function.
    4. At the top of this page, note that package `quote` is included in the `rsc.io/quote` module.

    You can use the pkg.go.dev site to find published modules whose packages have functions you can use in your own code. Packages are published in modules -- like `rsc.io/quote` -- where others can use them. Modules are improved with new versions over time, and you can upgrade your code to use the improved versions.

2.  In your Go code, import the `rsc.io/quote` package and add a call to its Go function.
    
    ```diff
     package main

     import "fmt"

    +import "rsc.io/quote"

     func main() {
    -    fmt.Println("Hello, Go!")
    +    fmt.Println(quote.Go())
     }
    ```

3.  Add new module requirements and sums.

    Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module. For more, see [Authenticating modules](https://go.dev/ref/mod#authenticating) in the Go Modules Reference.

    ```
    $ go mod tidy
    go: finding module for package rsc.io/quote
    go: found rsc.io/quote in rsc.io/quote v1.5.2
    ```

4.  Run your code to see the message generated by the function you're calling.

    ```
    $ go run .
    Don't communicate by sharing memory, share memory by communicating.
    ```

    Notice that your code calls the `Go` function, printing a clever message about communication.

    When you ran go `mod tidy`, it located and downloaded the `rsc.io/quote` module that contains the package you imported. By default, it downloaded the latest version -- v1.5.2.


## Part 2: Your first package

In this part, you'll write a small package with functions and use it from the `hello` program. Along the way you will get introduced to functions, error handling, unit testing, and compiling.

### Start a package that others can use

You'll start by creating a Go package. 

1.  Open a command prompt and cd to your home directory.

    ```
    cd
    ```

2.  Create a `${HOME}/hello/greetings` directory for your Go package source code.

    After you create this directory, you should have a `greetings` directory under the `hello` directory, like so:

    ```
    <home>/
    |-- hello/
        |-- greetings/
    ```

    For example, from your home directory use the following commands:

    ```
    cd hello
    mkdir greetings
    cd greetings
    ```

3.  In your text editor, create a file under the `greetings` directory in which to write your code and call it `greetings.go`.
   
4.  Paste the following code into your `greetings.go` file and save the file.

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

    This is the first code for your package. It returns a greeting to any caller that asks for one. You'll write code that calls this function in the next step.

    In this code, you:
    
    - Declare a greetings package to collect related functions.
    - Implement a Hello function to return the greeting.

      This function takes a name parameter whose type is `string`. The function also returns a `string`. In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name. For more about exported names, see [Exported names](https://go.dev/tour/basics/3) in the Go tour.
    
      <img src="assets/images/function-syntax.png" width="40%">

    - Declare a `message` variable to hold your greeting.
    
      In Go, the `:=` operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

      ```go
      var message string
      message = fmt.Sprintf("Hi, %v. Welcome!", name)
      ```

    - Use the `fmt` package's `Sprintf` function to create a greeting message. The first argument is a format string, and Sprintf substitutes the name parameter's value for the %v format verb. Inserting the value of the name parameter completes the greeting text.
    - Return the formatted greeting text to the caller.

### Call your code from another package

You'll write code you can execute as an application, and which calls the `Hello` function in the `greetings` package you just wrote.

1.  In your text editor, in the `hello` directory, create a file in which to write your code and call it `hello.go`.
2.  Write code to call the `Hello` function, then print the function's return value.
    
    To do that, paste the following code into `hello.go`.

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

    In this code, you:

    - Declare a main package. In Go, code executed as an application must be in a main package.
    - Import two packages: `example/hello/greetings` and the `fmt` package. This gives your code access to functions in those packages. Importing `example/hello/greetings` (the package you created earlier) gives you access to the `Hello` function. You also import `fmt`, with functions for handling input and output text (such as printing text to the console).
    - Get a greeting by calling the `greetings` package’s `Hello` function.

5.  At the command prompt in the `hello` directory, run your code to confirm that it works.

    ```
    $ go run .
    Hi, Gladys. Welcome!
    ```

### Return and handle an error

Handling errors is an essential feature of solid code. In this section, you'll add a bit of code to return an error from the greetings package, then handle it in the caller.

1.  There's no sense sending a greeting back if you don't know who to greet. Return an error to the caller if the name is empty. 

    In `greetings/greetings.go`, change your code to return an error with a message:

    ```diff
     package greetings

     import (
    +   "errors"
        "fmt"
     )

     // Hello returns a greeting for the named person.
    -func Hello(name string) (string) {
    +func Hello(name string) (string, error) {
    +    // If no name was given, return an error with a message.
    +    if name == "" {
    +        return "", errors.New("empty name")
    +    }
    +
        // If a name was received, return a value that embeds the name
        // in a greeting message.
        message := fmt.Sprintf("Hi, %v. Welcome!", name)
    -   return message
    +   return message, nil
    }
    ```

    In this code, you:

    - Change the function so that it returns two values: a `string` and an `error`. Your caller will check the second value to see if an error occurred. (Any Go function can return multiple values. For more, see [Effective Go](https://go.dev/doc/effective_go.html#multiple-returns).)
    - Import the Go standard library errors package so you can use its [`errors.New` function](https://pkg.go.dev/errors/#example-New).
    - Add an if statement to check for an invalid request (an empty string where the name should be) and return an error if the request is invalid. The `errors.New` function returns an error with your message inside.
    - Add `nil` (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.

2.  Change the `main()` function in `hello/hello.go` to handle both the value and the error now returned by the `Hello` function.

    ```diff
    package main

     import (
        "fmt"
    +   "log"

        "example/hello/greetings"
     )

     func main() {
    +   // Set properties of the predefined Logger, including
    +   // the log entry prefix and a flag to disable printing
    +   // the time, source file, and line number.
    +   log.SetPrefix("greetings: ")
    +   log.SetFlags(0)

        // Request a greeting message.
    -   message := greetings.Hello("Gladys")
    +   message, err := greetings.Hello("")
    +   // If an error was returned, print it to the console and
    +   // exit the program.
    +   if err != nil {
    +       log.Fatal(err)
    +   }

    +   // If no error was returned, print the returned message
    +   // to the console.
        fmt.Println(message)
     }
    ```

    In this code, you:

    - Configure the `log` package to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
    - Assign both of the `Hello` return values, including the `error`, to variables.
    - Change the `Hello` argument from Gladys’s name to an empty string, so you can try out your error-handling code.
    - Look for a non-nil `error` value. There's no sense continuing in this case.
    - Use the functions in the standard library's `log package` to output error information. If you get an error, you use the log package's [`Fatal` function](https://pkg.go.dev/log?tab=doc#Fatal) to print the error and stop the program.
    
3.  At the command line in the `hello` directory, run `hello.go` to confirm that the code works.

    Now that you're passing in an empty name, you'll get an error.

    ```
    $ go run .
    greetings: empty name
    exit status 1
    ```

That's common error handling in Go: Return an error as a value so the caller can check for it.

### Add a test

Now that you've gotten your code to a stable place, add a test. Testing your code during development can expose bugs that find their way in as you make changes. In this part, you add a test for the `Hello` function.

Go's built-in support for unit testing makes it easier to test as you go. Specifically, using naming conventions, Go's `testing` package, and the `go test` command, you can quickly write and execute tests.

1.  In the greetings directory, create a file called `greetings_test.go`.

    Ending a file's name with _test.go tells the `go test` command that this file contains test functions.

2.  In `greetings_test.go`, paste the following code and save the file.

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

    In this code, you:

    - Implement test functions in the same package as the code you're testing.
    - Create two test functions to test the `greetings.Hello` function. Test function names have the form `TestName`, where *Name* says something about the specific test. Also, test functions take a pointer to the `testing` package's [`testing.T` type](https://pkg.go.dev/testing/#T) as a parameter. You use this parameter's methods for reporting and logging from your test.
    - Implement two tests:
      - `TestHelloName` calls the `Hello` function, passing a name value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the `t` parameter's [`Fatalf` method](https://pkg.go.dev/testing/#T.Fatalf) to print a message to the console and end execution.
      - TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the t parameter's Fatalf method to print a message to the console and end execution.

3.  At the command line in the greetings directory, run the [`go test` command](https://go.dev/cmd/go/#hdr-Test_packages) to execute the test.

    The `go test` command executes test functions (whose names begin with `Test`) in test files (whose names end with _test.go). You can add the `-v` flag to get verbose output that lists all of the tests and their results.

    The tests should pass.

    ```
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

4.  Break the `greetings.Hello` function to view a failing test.

    The `TestHelloName` test function checks the return value for the name you specified as a `Hello` function parameter. To view a failing test result, change the `greetings.Hello` function in `greetings/greetings.go` so that it no longer includes the name.

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

5.  At the command line in the greetings directory, run go test to execute the test.

    This time, run `go test` without the `-v` flag. The output will include results for only the tests that failed, which can be useful when you have a lot of tests. The `TestHelloName` test should fail -- `TestHelloEmpty` still passes.

    ```
    $ go test
    --- FAIL: TestHelloName (0.00s)
        greetings_test.go:15: Hello("Gladys") = "Hi, Welcome!", <nil>, want match for `\bGladys\b`, nil
    FAIL
    exit status 1
    FAIL    example/hello/greetings   0.182s
    ```

### Compile and install the application

In this last topic, you'll learn a couple new go commands. While the `go run` command is a useful shortcut for compiling and running a program when you're making frequent changes, it doesn't generate a binary executable.

This topic introduces two additional commands for building code:

- The `go build` command compiles the packages, along with their dependencies, but it doesn't install the results.
- The `go install` command compiles and installs the packages.


1.  From the command line in the hello directory, run the `go build` command to compile the code into an executable.

    ```
    $ go build
    ```

2.  From the command line in the hello directory, run the new `hello` executable to confirm that the code works.

    Note that your result might differ depending on whether you changed your `greetings.go` code after testing it.

    ```
    $ ./hello
    Hi, Gladys. Welcome!
    ```

    You've compiled the application into an executable so you can run it. But to run it currently, your prompt needs either to be in the executable's directory, or to specify the executable's path.

    Next, you'll install the executable so you can run it without specifying its path.

3.  Discover the Go install path, where the go command will install the current package.

    You can discover the install path by running the go list command, as in the following example:

    ```
    $ go list -f '{{.Target}}'
    ```

    For example, the command's output might say `/home/gopher/bin/hello`, meaning that binaries are installed to /home/gopher/bin. You'll need this install directory in the next step.

4.  Add the Go install directory to your system's shell path.

    That way, you'll be able to run your program's executable without specifying where the executable is.

    ```
    $ export PATH=$PATH:/path/to/your/install/directory
    ```

    As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the [go env command](https://go.dev/cmd/go/#hdr-Print_Go_environment_information):

    ```
    $ go env -w GOBIN=/path/to/your/bin
    ```

5.  Once you've updated the shell path, run the `go install` command to compile and install the package.

    ```
    $ go install
    ```

    Run your application by simply typing its name. To make this interesting, open a new command prompt and run the `hello` executable name in some other directory.

    ```
    $ hello
    Hi, Gladys. Welcome!
    ```

## Exercises

### Return a random greeting

Change function `Hello` so that instead of returning a single greeting every time, it returns one of several predefined greeting messages.

To do this, you can use a Go slice. A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.

Add a small slice to contain three greeting messages, then have your code return one of the messages randomly. For more on slices, see [Go slices](https://blog.golang.org/slices-intro) in the Go blog. To learn about pseudo-random number generation in Go, refer to the [rand package](https://pkg.go.dev/math/rand).

### Return greetings for multiple people

Add a new function `Hellos` that returns greetings for multiple people.

This function should take a slice of names as input and return a map that associates each name with a personalized greeting message. Unlike the previous `Hello` function, which handled a single name and returned a single string, Hellos works with multiple names and returns a map from names to greetings.

To support this, you'll update the return type from a single `string` to a `map[string]string`, allowing you to generate and return a greeting for each name in the input slice.

```go
// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error)
```

## What's next

See [Effective Go](https://go.dev/doc/effective_go.html) for tips on writing clear, idiomatic Go code.

Take [A Tour of Go](https://go.dev/tour/) to learn the language proper.

Visit the [documentation page](https://go.dev/doc/#articles) for a set of in-depth articles about the Go language and its libraries and tools.