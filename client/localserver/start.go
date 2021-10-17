package localserver

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/static", "./static")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	r.LoadHTMLGlob(fmt.Sprintf("%s/*.html", exPath))

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	fmt.Printf("%s", r.Run(":4001"))
}
