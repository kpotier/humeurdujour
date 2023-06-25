package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kpotier/humeurdujour/pkg/api"
	store "github.com/kpotier/humeurdujour/pkg/store/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	e := gin.Default()

	api := api.API{Log: log.Default(), Store: store.NewStore(db)}
	a := e.Group("/api")
	api.Routes(a)

	e.Use(Static("/", "./static/dist"))

	e.Run()
}

func Static(prefix string, dir string) gin.HandlerFunc {
	d := static.LocalFile(dir, false)
	fs := http.FileServer(d)
	if prefix != "" {
		fs = http.StripPrefix(prefix, fs)
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Next()
		} else if d.Exists(prefix, c.Request.URL.Path) {
			fs.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.Request.URL.Path = "/"
			fs.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
