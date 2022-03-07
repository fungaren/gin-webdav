package webdav

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	wd "golang.org/x/net/webdav"
)

func Serve(prefix string, rootDir string,
	logger func(req *http.Request, err error)) gin.HandlerFunc {
	w := wd.Handler{
		Prefix:     prefix,
		FileSystem: wd.Dir(rootDir),
		LockSystem: wd.NewMemLS(),
		Logger:     logger,
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, w.Prefix) {
			c.Status(200) // 200 by default, which maybe override later
			w.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
