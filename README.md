# GoAdmin+Vue Example

A example show how to develop GoAdmin with vue.

一个例子展示如何利用GoAdmin结合vue的前后端开发流程。

backend文件夹下是所有后端开发文件，frontend文件夹下是所有前端开发文件。

![](http://quick.go-admin.cn/docs/vue-goadmin.png)


## 步骤一

后端编译二进制执行文件给前端。前端将文件放置于文件夹下。当前代码依赖最新master分支的GoAdmin。使用go mod方式，执行以下命令：

```bash
go build .
cp backend ./../frontend
```

## 步骤二

前端在文件夹下执行：

```bash
cd frontend
./backend --debug=true
```

打开：[http://localhost:9033/vue](http://localhost:9033/vue)，可以看到对应页面。

程序会自动监听```./src/src```下文件变化，有变化则执行```npm run build```命令进行编译。

以上。