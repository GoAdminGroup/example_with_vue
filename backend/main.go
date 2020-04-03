package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                    // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		uploadDir, assetDir, srcDir, port, index string
		debug, watchMode                         bool
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&watchMode, "watch", true, "watch mode")
	flag.StringVar(&uploadDir, "upload", "./uploads", "upload dir")
	flag.StringVar(&assetDir, "assets", "./dist/static", "assets dist dir")
	flag.StringVar(&index, "index", "./dist/index.html", "index html path")
	flag.StringVar(&srcDir, "src", "./src/src", "frontend src path")
	flag.StringVar(&port, "port", "9033", "http listen port")
	flag.Parse()

	r := gin.Default()

	if !debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON("./config.json").
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", uploadDir)
	r.Static("/static", assetDir)

	eng.HTMLFile("GET", "/admin/vue", index, map[string]interface{}{})
	eng.HTMLFile("GET", "/admin/vue/*any", index, map[string]interface{}{})

	if watchMode {
		go watch(srcDir)
	}

	_ = r.Run(":" + port)
}

func watch(watchDIR string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = watcher.Close()
	}()

	done := make(chan bool)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("watch error: ", err)
			}
		}()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if !strings.Contains(event.Name, "___jb_tmp__") &&
					(event.Op&fsnotify.Write == fsnotify.Write ||
						event.Op&fsnotify.Create == fsnotify.Create) {
					cmd := exec.Command("npm", "--prefix", "./src", "run", "build")
					err = cmd.Run()
					fmt.Println("building...")
					if err != nil {
						fmt.Println("build error", err)
					} else {
						fmt.Println("build success")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	_ = filepath.Walk(watchDIR, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			checkErr(watcher.Add(path))
			fmt.Println("Monitoring Dir: ", path)
		}
		return nil
	})

	<-done
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
