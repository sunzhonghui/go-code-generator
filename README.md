## Golang 通用代码生成器 go-code-generator

#### v0.0.1

### 简介

> 采用Fyne 编写的Gui客户端，跨平台编译
>
> 根据项目架构连接数据库动态生成go代码
>
> 可以按照自己项目结构生成不同的目录结构
>
> 现在只支持MySQL

### 依赖组件

- gorm-v2 
- gin
- gin-swagger

### 生成的结构

现在生成的类型

- model
- mapper
- service
- api
- router

生成的目录结构 会生成在当前程序运行目录 gen-code/ 下

------router 

------project

----模块1

--model

--mapper

--service

--api

----模块2

--model

--mapper

--service

--api

……

### 动态模板

可以根据不同架构自行编辑生成模板文件

模板文件会在项目启动的时候

自动装载到当前程序运行目录 temp/ 下



