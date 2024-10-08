## 后台模版

初始化数据库 go-vue-amis.sql

默认账户
admin 123456

访问 http://127.0.0.1:8080/web

前端
https://github.com/go-home-admin/vue-amis-admin

安装gnumake 注意安装完整版本
#https://jaist.dl.sourceforge.net/project/getgnuwin32/GetGnuWin32_legacy_install_archive.zip?viasf=1   --header "Cookie:_ga=GA1.1.1582004290.1724493212; _ga_1H226E4E4L=GS1.1.1725443719.2.1.1725443916.0.0.0;"    --header "Referer:https://sourceforge.net/"    --header "User-Agent:Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0"

#env.copy 改名为.env
```shell
cp .env.copy .env
```
安装依赖
```shell 
go install github.com/golang/protobuf/protoc-gen-go@latest		# proto 工具链, 生成go代码插件
go install github.com/go-home-admin/toolset@latest
```

生成代码
```shell
make gen
```
开始运行
```shell
make dev
```



```shell 
toolset make:orm
toolset make:protoc
toolset make:route
toolset make:bean
toolset make:swagger
gen openapi.json to /Users/lv/Desktop/github.com/go-admin/web/swagger.json
```

# set GOHOSTARCH amd64
```shell
go env -w GOARCH=amd64
```
 