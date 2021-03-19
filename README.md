## 快速创建一个新项目

###  GO应用项目基本布局
[Standard Go Project Layout](https://github.com/golang-standards/project-layout)

#### 关键几个包：
* cmd：本项目的主干, 通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西。

* internal：私有应用程序和库代码，并不局限于顶级 internal 目录。在项目树的任何级别上都可以有多个内部目录

* pkg：外部应用程序可以使用的库代码。

* vendor：应用程序依赖项

* config：配置文件模板或默认配置 

* scripts：执行各种构建、安装、分析等操作的脚本。这些脚本保持了根级别的 Makefile 变得小而简单

* build：打包和持续集成,将你的 CI (travis、circle、drone)配置和脚本放在 /build/ci 目录中；docker相关配置放在/build/docker；k8s的相关配置放在build/k8s

* githooks：截获本机 git commit 命令，执行定制脚本

* test：额外的外部测试应用程序和测试数据


当然这只是一种规范，其他的包名以及层级结构也是灵活的。该项目会严格规范来命名和创建


### Git提交规范

// TODO 规范说明 （husky工具 ｜ githook加脚本验证）
* commit-msg check
* gofmt format code

### golangci-lint
[golangci-lint](https://github.com/golangci/golangci-lint)

// TODO 


### CICD

// TODO（通过脚手架生成dockerfile k8s的yaml文件）







