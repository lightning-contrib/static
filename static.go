package static

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-labx/lightning"
)

// config struct holds the root and prefix values
type config struct {
	root   string
	prefix string
}

// Options is a function that takes a pointer to a config struct
type Options func(*config)

// WithRoot is a function that takes a string and returns an Options function
func WithRoot(root string) Options {
	return func(cfg *config) {
		cfg.root = root
	}
}

// WithPrefix is a function that takes a string and returns an Options function
func WithPrefix(prefix string) Options {
	return func(cfg *config) {
		cfg.prefix = prefix
	}
}

// Default returns a new instance of the middleware with default options
func Default() lightning.Middleware {
	return New()
}

// New returns a new instance of the middleware with the provided options
func New(options ...Options) lightning.Middleware {
	// set default values for the config struct
	cfg := &config{
		root:   "./public",
		prefix: "/static/",
	}

	// apply any provided options to the config struct
	for _, option := range options {
		option(cfg)
	}

	// get the path of the executable file
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	return func(ctx *lightning.Context) {
		path := ctx.Path
		// if the request method is not GET or the path does not start with the prefix, skip this middleware
		if ctx.Method != http.MethodGet || !strings.HasPrefix(path, cfg.prefix) {
			ctx.Next()
			return
		}

		// get the full file path by joining the root and the path after the prefix
		fullFilePath := filepath.Join(exPath, cfg.root, strings.TrimPrefix(path, cfg.prefix))
		// if the file exists, serve it with a 200 status code
		if _, err := os.Stat(fullFilePath); !os.IsNotExist(err) {
			ctx.SkipFlush()
			ctx.SetStatus(http.StatusOK)
			http.ServeFile(ctx.Res, ctx.Req, fullFilePath)
		} else {
			ctx.Text(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}
	}
}
