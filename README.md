# GoAdmin+Vue Example

A example show how to develop GoAdmin with vue.

这是一个例子展示如何利用GoAdmin结合vue的前后端开发流程。

backend文件夹下是所有后端开发文件，frontend文件夹下是所有前端开发文件。

**注：本例子依赖sqlite，如果你使用windows，请先下载安装gcc。**

![](http://quick.go-admin.cn/docs/vue-goadmin-2.png)

## 前端UI框架

选取目前主流的三个UI vue Admin框架作为例子：

- [adminlte](https://github.com/devjin0617/vue2-admin-lte)
- [element ui](https://github.com/PanJiaChen/vue-element-admin)
- [ant design](https://github.com/iczer/vue-antd-admin)

## 开发流程

### 步骤一

后端开发人员编译二进制执行文件给前端开发人员。前端开发人员将二进制可执行文件放置于前端文件夹下。

当前仓库代码依赖最新master分支的GoAdmin。使用go mod方式，执行以下命令：

```bash
git clone https://github.com/GoAdminGroup/goadmin-vue-example.git
cd goadmin-vue-example/backend

# build binary file
GO111MODULE=on go build -o ./backend .
# send the binary file to all frontend developers
cp backend ./../frontend
```

### 步骤二

前端在前端文件夹下执行：

```bash
cd frontend
# install dependencies
npm install --registry=https://registry.npm.taobao.org --prefix ./src/element
# build vue project
NODE_ENV=production npm --prefix ./src/element run build
# use the backend binary file to serve instead of node
NODE_ENV=production ./backend  --debug=true --theme=sword --ui_theme=element 
```

更多命令，请看 Makefile。

打开：[http://localhost:9033/admin/login](http://localhost:9033/admin/login)

账号：admin 密码：admin，登录后可以看到对应vue页面。

程序会自动监听```./src/src```下文件变化，有变化则执行```npm --prefix ./src run build```命令进行编译。

执行：```./backend --help``` 可以看到对应帮助选项信息。

以上。