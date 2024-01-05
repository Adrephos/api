package middlewares

import (
	"bytes"
	"strings"

	"github.com/Adrephos/api/utils"
	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func CacheAdd(c *gin.Context) {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w

	cache := utils.GetCache()
	key := strings.Clone(c.Request.URL.String())
	key = strings.TrimRight(key, "/")

	val, found := cache.Get(key)
	if found {
		c.Header("Content-Type", val.(utils.CacheEntry).ContentType)
		c.Writer.Write(val.(utils.CacheEntry).Body)
		c.Status(val.(utils.CacheEntry).StatusCode)
		c.Abort()
		return
	}

	c.Next()

	if c.Writer.Status() == 200 {
		cache.Set(
			key,
			utils.CacheEntry{
				Body:        w.body.Bytes(),
				StatusCode:  c.Writer.Status(),
				ContentType: strings.Clone(c.Writer.Header().Get("Content-Type")),
			},
			0,
		)
	}
	return
}
