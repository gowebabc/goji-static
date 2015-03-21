# goji-static

Package `goji-static` provides a middleware to handle static asset files. It also allows to leverage browser cache by using cache control headers and ETags.

It also allows access to `/favicon.ico`, which can be stored along with the `main.go` project file. Commonly, many browsers will try to fetch a favicon from the domain root if they can't find the favicon tag in the HTML.

The cache control header is set up as this: `Cache-Control: public, max-age=31536000` which is the recommended way to handle files in the browser cache for almost a year, which is recommended by Google.

### Usage example

```go
package main

import (
    "github.com/theosomefactory/goji-static"
    "github.com/zenazn/goji"
)

func main() {
	// Serves the static files from the `assets` folder
    goji.Use(static.Static("assets"))
    goji.Get("/", myHandler)
    goji.Serve()
}
```
