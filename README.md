## Golang 通用代码生成器 go-code-generator

#### v0.0.1

![首页](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/main.png)

### 简介

> 采用Fyne 编写的Gui客户端，跨平台编译
>
> 根据项目架构连接数据库动态生成go代码
>
> 可以按照自己项目结构生成不同的目录结构
>
> 现在只支持MySQL

### 依赖组件

- gorm-v2 [https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
- gin-gonic [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- swaggo [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

### 自动加载配置

项目启动的时候会自动创建配置文件夹

自动加载默认配置文件到 /resource 下

font/ 字体文件
temp/ 模板文件
conf.yaml 默认配置文件

### 配置




### 打包

fyne package -os windows -icon idmisstx.png
……




