# Go Url Shortener
This exercise is not actually a shortener such as bitly, it's an exercise to show the use of `http.Handler`(s) and middleware.

- Generic HTTP Server
- HTTP Mux implementation (request multiplexer)
- Middleware that returns http.Handler
- Read in JSON file
- Parse JSON byte array to pointer(map)

### Getting started

```
// pull
git clone https://github.com/mmason33/go-url-shortener.git

// run
go run main.go
```

### Optional Flag
```
// flag -redirects
// defaults to redirects.json
go run main.go -redirects=./test.json

// view flags
go run main.go -h

```