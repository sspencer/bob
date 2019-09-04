package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (s *server) routes() {
	// public routes
	r := s.engine.Group("/v1")
	r.GET("/version", versionHandler)
}

// GET /v1/version
func versionHandler(c *gin.Context) {
	v := "0.0.0 (unset)"
	if Version != "" {
		v = Version
	}

	c.JSON(200, gin.H{
		"name":    "bob",
		"uptime":  time.Duration(time.Since(uptime)).String(),
		"version": v,
	})
}
