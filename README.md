# GoAdmin+Vue Example

A example show how to develop GoAdmin with vue.

一个例子展示如何利用GoAdmin结合vue的前后端开发流程。

backend文件夹下是所有后端开发文件，frontend文件夹下是所有前端开发文件。

![](http://quick.go-admin.cn/docs/vue-goadmin.png)


## 步骤一

后端开发人员编译二进制执行文件给前端开发人员。前端开发人员将二进制可执行文件放置于前端文件夹下。

当前仓库代码依赖最新master分支的GoAdmin。使用go mod方式，执行以下命令：

```bash
# build binary file
GO111MODULE=on go build -o ./backend .
# give the binary file to all frontend developers
cp backend ./../frontend
```

## 步骤二

前端在前端文件夹下执行：

```bash
cd frontend
# build vue
npm --prefix ./src run build
# use the backend binary file to serve instead of node
./backend --debug=true
```

打开：[http://localhost:9033/admin/login](http://localhost:9033/admin/login)，登录后可以看到对应vue页面。

程序会自动监听```./src/src```下文件变化，有变化则执行```npm --prefix ./src run build```命令进行编译。

执行：```./backend --help``` 可以看到对应帮助选项信息。

以上。