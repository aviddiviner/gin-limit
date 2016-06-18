package limit

import (
	"github.com/gin-gonic/gin"
)

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire() // before request
		c.Next()
		release() // after request
	}
}

// Simple as that. Now you know :)
