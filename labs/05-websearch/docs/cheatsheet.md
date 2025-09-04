# Go Cheatsheet

## Implementing HTTP Handlers

When implementing HTTP handlers in Go, you’ll use the signature:

```go
func(w http.ResponseWriter, r *http.Request)
```

This guide summarizes key components you may need when working with `http.Request`, `http.ResponseWriter`, URLs, query parameters, HTML files, and JSON responses.

### `r *http.Request`: Accessing Request Data

Get query parameters from the URL:
- API Reference: `r.URL.Query()`
- Returns a url.Values type (i.e., a map of string → []string)

Get a specific query parameter:
- API Reference: url.Values.Get(key string)
- Retrieves the first value for a named query parameter (e.g., "q")

### `w http.ResponseWriter`: Writing Responses

Set the response header:
- API Reference: `w.Header().Set(key, value)`
- Common headers: `"Content-Type": "application/json"` or `"text/html"`

Write an error response:
- API Reference: `http.Error(w, msg, statusCode)`
- Sends a plain-text error message with the given HTTP status code

Write a response body:
- API Reference: `w.Write([]byte)`
- Use this to send HTML or plain-text output

### Processing Strings
Use the [strings](https://pkg.go.dev/strings) package to:
- Convert to lowercase: `strings.ToLower(string)`
- Split input by whitespace: `strings.Fields(string)`

These are useful for cleaning and tokenizing user input, e.g., search queries.

### Reading Static Files
Use the [os](https://pkg.go.dev/os) package to read the contents of a file (e.g., an HTML page):
- API Reference: `os.ReadFile(path)`
- Reads the entire file into a byte slice ([]byte)
- Returns (content []byte, err error)

### Encoding Results as JSON
Use the [json](https://pkg.go.dev/encoding/json) package to send structured data (e.g., search results) in the HTTP response:
- API Reference: `json.NewEncoder(w).Encode(v)`
- Encodes a Go value as JSON and writes it to the response writer

Before encoding JSON:

```go
w.Header().Set("Content-Type", "application/json")
```

### Tips
- Always check for errors (reading files, scanning, encoding, etc.)
- Normalize user input: lowercase, trim, split
- Return clear status codes:
  - 400 for bad input
  - 500 for internal errors
- For HTML content, set `"Content-Type"` to `"text/html"`
- For JSON APIs, set `"Content-Type"` to `"application/json"`

## Reading and Scanning Text Files Line-by-Line
Use the [os](https://pkg.go.dev/os) and [bufio](https://pkg.go.dev/bufio) packages to read structured text files.

### Open the file:
- API Reference: `os.Open(filename)`
- Returns `*os.File` and error

### Scan line-by-line:
- API Reference: `bufio.NewScanner(file)`
- Use `scanner.Scan()` in a loop
- Access each line with `scanner.Text()`
- Example:  
  ```go
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    // process line
  }
  ```

### Close the file when done:
- API Reference: `file.Close()`
- Important to free resources
- Tip: use defer `file.Close()` right after opening

Check for scanner errors:
- API Reference: `scanner.Err()`

### Tips
- Use defer `file.Close()` after opening files
