package localserver

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func StartServer(openUrl string) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.Static("/static", "./static")

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	r.Use(static.Serve("/", static.LocalFile(fmt.Sprintf("%s/site", exPath), true)))

	// r.LoadHTMLGlob(fmt.Sprintf("%s/site/*.html", exPath))

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })
	if openUrl != "" {
		go func() {
			time.Sleep(1500 * time.Millisecond)
			open(openUrl)
		}()
	}

	fmt.Printf("%s", r.Run(":4001"))
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
