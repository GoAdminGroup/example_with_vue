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
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		uploadDir, assetDir, srcDir, port, index, watchIgnore, configPath string
		debug, watchMode, quiet                                           bool
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&watchMode, "watch", true, "watch mode")
	flag.BoolVar(&quiet, "quiet", false, "disable log")
	flag.StringVar(&watchIgnore, "ignore", "___jb", "watch ignore files, use comma split")
	flag.StringVar(&uploadDir, "upload", "./uploads", "upload dir")
	flag.StringVar(&assetDir, "assets", "./dist/static", "assets dist dir")
	flag.StringVar(&configPath, "config", "./config.json", "config path")
	flag.StringVar(&index, "index", "./dist/index.html", "index html path")
	flag.StringVar(&srcDir, "src", "./src/src", "frontend src path")
	flag.StringVar(&port, "port", "9033", "http listen port")
	flag.Parse()

	if !debug || quiet {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	r := gin.New()

	eng := engine.Default()

	cfg := config.ReadFromJson(configPath)

	if quiet {
		cfg.Debug = false
		cfg.AccessLogOff = true
		cfg.ErrorLogOff = true
		cfg.InfoLogOff = true
	}

	if err := eng.AddConfig(cfg).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", uploadDir)
	r.Static("/static", assetDir)

	eng.HTMLFile("GET", "/admin/vue", index, map[string]interface{}{})
	eng.HTMLFile("GET", "/admin/vue/*any", index, map[string]interface{}{})

	if watchMode {
		go watch(srcDir, watchIgnore, quiet)
	}

	_ = r.Run(":" + port)
}

func watch(dir, ignores string, logOff bool) {

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
				printOut(logOff, "watch recover error: ", err)
			}
		}()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					printOut(logOff, "watch events error")
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
					printOut(logOff, path.Base(event.Name)+" change")
					printOut(logOff, "building....")
					cmd := exec.Command("npm", "--prefix", "./src", "run", "build")
					err = cmd.Run()
					if err != nil {
						printOut(logOff, "build error", err)
					} else {
						printOut(logOff, "build success")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					printOut(logOff, "watch errors")
					break
				}
				printOut(logOff, "error:", err)
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
			printOut(logOff, "Monitoring Dir: ", p)
		}
		return nil
	})

	<-done
}

func printOut(logOff bool, a ...interface{}) {
	if !logOff {
		fmt.Println(a...)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
