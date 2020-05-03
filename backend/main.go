package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                    // ui theme
	_ "github.com/GoAdminGroup/themes/sword"                       // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		uploadDir, assetDir, srcDir, port, index, watchIgnore, configPath, uitheme, theme string
		debug, watchMode, quiet                                                           bool
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&watchMode, "watch", true, "watch mode")
	flag.BoolVar(&quiet, "quiet", false, "disable log")
	flag.StringVar(&watchIgnore, "ignore", "___jb", "watch ignore files, use comma split")
	flag.StringVar(&uploadDir, "upload", "./uploads", "upload dir")
	flag.StringVar(&assetDir, "assets", "./dist/static", "assets dist dir")
	flag.StringVar(&configPath, "config", "./config.json", "config path")
	flag.StringVar(&index, "index", "./dist/index.html", "index html path")
	flag.StringVar(&srcDir, "src", "./src/", "frontend src path")
	flag.StringVar(&port, "port", "9033", "http listen port")
	flag.StringVar(&uitheme, "ui_theme", "adminlte", "ui theme")
	flag.StringVar(&theme, "theme", "adminlte", "GoAdmin theme")
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

	fmt.Println("cfg.Theme", cfg.Theme)

	if err := eng.AddConfig(cfg).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", uploadDir)
	r.Static("/static", assetDir)

	eng.HTMLFile("GET", "/admin/vue", index, map[string]interface{}{})
	eng.HTMLFile("GET", "/admin/vue/*any", index, map[string]interface{}{})

	if watchMode {
		go watch(srcDir+uitheme+"/src", watchIgnore, uitheme, quiet)
	}

	r.POST("/mock-server/user/info", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, `{"code":200,"msg":"success","data":{"roles":["admin"],"userName":"admin"}}`)
	})

	r.POST("/mock-server/changeLog/getList", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, `{"code":200,"msg":"success","totalCount":999,"data":[{"content":"在github上获得了第一个star，感恩一位名叫Bequiet2014的github用户","timestamp":"2020-03-23"},{"content":"增加更换主题功能","timestamp":"2020-04-10"},{"content":"大幅精简代码","timestamp":"2020-04-14"},{"content":"修复群友反馈的bug","timestamp":"2020-04-16"},{"content":"剔除maptalks","timestamp":"2020-04-17"},{"content":"换行符统一修改为lf 支持苹果 linux windows协同开发 强制开启最严格eslint规则 不要哭 严格是有好处的","timestamp":"2020-04-17"},{"content":"彻底完成手机端适配,记录这一天熬夜到了晚上三点","timestamp":"2020-04-18"},{"content":"删除babel-polyfill 提高打包速度 减少压缩体积（放弃ie是这个项目做出的最伟大的决定）","timestamp":"2020-04-18"},{"content":"源码精简至800k","timestamp":"2020-04-19"},{"content":"添加视频播放器组件","timestamp":"2020-04-20"},{"content":"修复路由懒加载 完善主题配色","timestamp":"2020-04-22"},{"content":"修复全局axios拦截 加快动画展示效果 修改登录页样式","timestamp":"2020-04-24"},{"content":"简化权限与登录逻辑 更新mock-server","timestamp":"2020-04-25"},{"content":"优化登录退出逻辑 代码更清晰 退出不再重载网页 改为重载路由形式","timestamp":"2020-04-26"},{"content":"无端的指责只会让我更加努力 修复sidebar 简化permission","timestamp":"2020-04-28"},{"content":"又是一个深夜 实现了表格增删改查的一键生成","timestamp":"2020-04-30"},{"content":"大幅优化tagsview标签动画","timestamp":"2020-05-02"},{"content":"三种图标组件实现mock模拟分页","timestamp":"2020-05-03"}]}`)
	})

	_ = r.Run(":" + port)
}

func watch(dir, ignores, theme string, logOff bool) {

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
					cmd := exec.Command("npm", "--prefix", "./src/"+theme, "run", "build")
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
