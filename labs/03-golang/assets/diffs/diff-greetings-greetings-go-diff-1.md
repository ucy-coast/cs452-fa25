#### Diff `greetings/greetings.go`

```diff
 package greetings

 import (
     "errors"
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
-    return message
+    return message, nil
 }
```

#### New `greetings/greetings.go`

```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```


#### Old `greetings/greetings.go`

```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string) {
    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```
