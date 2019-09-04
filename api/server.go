package api

import (
	"database/sql"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

const (
	Version = "0.0.1"
)

var (
	uptime time.Time
)

func init() {
	uptime = time.Now()
}

type server struct {
	engine   *gin.Engine
	database *sql.DB
}

func Server(database *sql.DB) *gin.Engine {

	engine := gin.Default() // default has Logger/Recovery middleware

	// Set a lower memory limit for multipart forms (defaults to 32 MiB)
	engine.MaxMultipartMemory = 2 << 20 // 2 MiB

	// All routes may return GZIP responses:
	//     curl -H "Accept-Encoding: gzip" ...
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	srv := server{
		engine:   engine,
		database: database,
	}

	srv.routes()

	return srv.engine
}
