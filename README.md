# WebDAV server for gin-gonic

## Usage

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"webdav/webdav"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil) // No proxy used
	r.Use(webdav.Serve(
		"/fs/",
		"/home/fang/test/",
		func(c *gin.Context) bool {
			// Anyone can access webdav server
			return true
		},
		func(req *http.Request, err error) {
			log.Println(req.URL.Path, err)
		},
	))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "")
	})
	r.Run(":7991")
}
```

## Test it

```bash
sudo apt install cadaver
cadaver http://127.0.0.1:7991/fs
```
