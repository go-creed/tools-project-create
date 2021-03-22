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


其他目录：

* assets: 静态文件、图片等 
  
* tools: 工具类相关

    ...

当然这只是一种规范，其他的包名以及层级结构也是灵活的。该项目会严格规范来命名和创建


### Git提交规范
commit message格式

`<type>(<scope>): <subject>`

**type(必须)**

用于说明git commit的类别，只允许使用下面的标识。

* feat：新功能（feature）。

* fix/to：修复bug，可以是QA发现的BUG，也可以是研发自己发现的BUG。

  * fix：产生diff并自动修复此问题。适合于一次提交直接修复问题
  * to：只产生diff不自动修复此问题。适合于多次提交。最终修复问题提交时使用fix

* docs：文档（documentation）。

* style：格式（不影响代码运行的变动）。

* refactor：重构（即不是新增功能，也不是修改bug的代码变动）。

* perf：优化相关，比如提升性能、体验。

* test：增加测试。

* chore：构建过程或辅助工具的变动。

* revert：回滚到上一个版本。

* merge：代码合并。

* sync：同步主线或分支的Bug。

**scope(可选)**

scope用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。

**subject(必须)**

subject是commit目的的简短描述，不超过50个字符。

如果英文表达不是很好，建议使用中文（用中文描述问题能更清楚一些）


### golangci-lint
[golangci-lint 静态代码检测工具](https://github.com/golangci/golangci-lint)


### 利用git hook 对代码和提交进行校验

项目生成时，会在githooks目录下会生成commit-msg 和 pre.commit文件定义了commit的规范检查和触发golangci-lint进行代码检查，主要通过软连接的方式将githook下的文件链接到该项目/.git/hook目录下


### CICD
项目创建成功，编写了






