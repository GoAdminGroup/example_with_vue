package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                    // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		uploadDir, assetDir, srcDir, port, index, watchIgnore string
		debug, watchMode                                      bool
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&watchMode, "watch", true, "watch mode")
	flag.StringVar(&watchIgnore, "ignore", "___jb", "watch ignore files, use comma split")
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

	if err := eng.AddConfigFromJSON("./config.json").
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", uploadDir)
	r.Static("/static", assetDir)

	eng.HTMLFile("GET", "/admin/vue", index, map[string]interface{}{})
	eng.HTMLFile("GET", "/admin/vue/*any", index, map[string]interface{}{})

	if watchMode {
		go watch(srcDir, watchIgnore)
	}

	_ = r.Run(":" + port)
}

func watch(dir, ignores string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = watcher.Close()
	}()

	igFiles := strings.Split(ignores, ",")

	done := make(chan bool)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("watch recover error: ", err)
			}
		}()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("watch events error")
					break
				}

				ignore := false

				// ignore jetbrains files change
				for _, f := range igFiles {
					if strings.Contains(event.Name, f) {
						ignore = true
						break
					}
				}

				if !ignore && (event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) {
					fmt.Println(path.Base(event.Name) + " change")
					fmt.Println("building....")
					cmd := exec.Command("npm", "--prefix", "./src", "run", "build")
					err = cmd.Run()
					if err != nil {
						fmt.Println("build error", err)
					} else {
						fmt.Println("build success")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					fmt.Println("watch errors")
					break
				}
				fmt.Println("error:", err)
			}
		}
	}()

	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			p, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			checkErr(watcher.Add(p))
			fmt.Println("Monitoring Dir: ", p)
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
