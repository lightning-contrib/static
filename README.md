# static

The `static` middleware is a Go package that provides a middleware for serving static files over HTTP. 

## Installation

To use the `static` middleware in your Go project, you can install it using go get:

```bash
go get github.com/go-labx/lightning-contrib/static
```

## Usage

To use the static middleware in your Go web application, you can import it and add it to your middleware stack:

```go
package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/static"
)

func main() {
	app := lightning.DefaultApp()

	// Serve static files from the "./public" directory under the "/static" URL prefix
    app.Use(static.New(
        static.WithRoot("./public"), 
        static.WithPrefix("/static")),
    )

	app.Get("/ping", func(ctx *lightning.Context) {
		ctx.Text(200, "hello world")
	})

	app.Run()
}
```

The `static.New` function returns a new instance of the middleware with the provided options. The `static.WithRoot` and `static.WithPrefix` functions are options that can be used to configure the root directory and URL prefix for serving static files.

By default, the `static` middleware serves files from the "./public" directory under the "/static" URL prefix. You can change these defaults by providing your own options to the `static.New` function.

## How it works

The `static` middleware works by intercepting HTTP requests and checking if the requested URL matches the configured URL prefix. If it does, the middleware serves the corresponding file from the configured root directory.

For example, if the middleware is configured to serve files from the "./public" directory under the "/static" URL prefix, a request for "/static/css/style.css" would be served from the file "./public/css/style.css".

If the requested file does not exist, the middleware returns a 404 Not Found error.

## API Documentation

For detailed API documentation and usage examples, please refer to the [documentation](https://pkg.go.dev/github.com/lightning-contrib/static).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](https://github.com/lightning-contrib/static/blob/main/CONTRIBUTING.md) for more information.

## License

This middleware is licensed under the [MIT License](https://raw.githubusercontent.com/lightning-contrib/static/main/LICENSE).
