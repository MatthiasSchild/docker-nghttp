package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-schild/logging"
	"io/ioutil"
	"os"
	"path"
)

var (
	webDir string
)

func init() {
	webDir = os.Getenv("WEB_DIR")
	if webDir == "" {
		webDir = "./web"
	}
}

func main() {
	r := gin.Default()
	r.Static("/", webDir)
	r.NoRoute(func(c *gin.Context) {
		data, _ := ioutil.ReadFile(path.Join(webDir, "index.html"))
		_, _ = c.Writer.Write(data)
	})
	logging.LogErr(r.Run())
}
