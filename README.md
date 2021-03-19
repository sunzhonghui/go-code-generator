## Golang 通用代码生成器 go-code-generator

#### v1.0.0

![alt 首页](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/mainv1.png)

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
![alt 启动前](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/folder-1.png)

```
项目启动的时候会自动创建配置文件夹

自动加载默认配置文件到 /resource 下

font/ 字体文件
temp/ 模板文件
conf.yaml 默认配置文件
```
![alt 启动后2](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/folder-2.png)

![alt 启动后3](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/folder-3.png)

![alt 启动后4](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/folder-4.png)


### 设置-数据库配置

![alt 数据库配置](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/set-database.png)

```
    输入对应的数据库配置

    点击测试，成功

    会自动保存到配置文件，下次打开程序自动读取数据库配置
```

### 设置-项目配置

![alt 项目配置](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/set-project.png)

##### 项目名称
```
    只用于展示
```
##### mod名称
```
    填写项目go.mod中 module名称，用于生成代码时import的前缀
```
    
##### 模块缩写
```
    模块的名称，用于生成代码的层级文件夹，跟import路径

```
##### 路由前缀
```
    生成接口的前缀

```

### 生成器-代码生成
![alt 代码生成](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/code-gen.png)

![alt 代码生成](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/code-gen-2.png)

```
    如果数据库测试成功，切换到“代码生成”页面，会自动列出来数据库所有的表
    点击选择表
    可以数据搜索的表名，模糊搜索，自动选择搜索到的第一张表
    点击生成就会按照配置的生成golang 代码
```
![alt 代码生成](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/code-gen-3.png)


     
### 例子

使用我自己搭建的简单 Golang 框架

目录结构

![alt 目录结构](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/demo-1.png)

我的配置项是这样的

![alt 配置](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/demo-2.png)

选择test表，点击生成，生成对应的代码

![alt 代码](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/demo-3.png)

![alt 代码](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/demo-4.png)

把代码拷贝到自己的项目，运行一下代码

已经有新的接口了

![alt 代码](https://raw.githubusercontent.com/sunzhonghui/go-code-generator/master/resource/images/demo-5.png)






### 打包

fyne package -os windows -icon idmisstx.png
……




