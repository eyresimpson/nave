# Nave Runner（Ambassador）

> 🍵 Creator: Noah Jones
> 
> ⌚️ Update: 2023 - 08 - 27
> 
> 🍃 Licence: MIT

> 🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
> 
> 🚌 🚗 Don't forget the light is always there 🚗 🚌
> 
> 🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳

> Nave Runner是一款类似于 Jvm 的执行器，和Jvm执行Class不同，Nave Runner主要用执行各种符合规范的蓝图文件，理论上任何文件均可以作为蓝图执行，只要可以将文件或数据表中的数据结构转换为符合规范的蓝图结构即可，Nave应用天然跨平台（和Jvm跨平台的方式一样），是一种解释执行的可视化编程
> 
> 在编写Nave时，我坚持以灵活的业务语言为基准，它的目标人群应该是业务人员和部分简单的场景，而不是对性能有高要求或极复杂的定制化的开发者
> 
> 我将Nave定义为一个简单低门槛的编程语言，而不是再创造一个Python或Java语言，我想编写一个简易轻松，快速高效的执行语言，我认为它应该适用于服务量少，业务结构简单，计划投入较少精力的场景，如果你在寻找一个处理复杂业务的底层语言，仍然简易使用Java，毕竟它具有成熟的框架，可以更好地处理复杂的任务
>
> Nave Ambassador 相当于全功能版的Runner，是Runner的一种运行模式，通过配置文件开启，默认关闭，它不再严格限制各种安全性操作，您可以直接通过API，实时的修改和调用各个流甚至直接执行Worker
>
> 系统通过插件方式扩展功能，用户可以通过类似于游戏 mod 的方式，快速的添加各种需要的功能库（这部分内容主要由设计器实现，执行器主要用于下载需要的依赖文件（windows下为dll、Linux、Macos下为so
>
> Nave具有一款跨多平台的设计系统（通过 Electron + Vue3实现），可以快速以图形化方式搭建蓝图

## 执行原理

Nave可以解析Json文件，并根据规范，开启两种不同类型的业务流：

- 接口业务流（接口）
- 定时业务流（定时任务）

当满足业务流的条件触发时，Nave会通过协程的方式调用流，此处以接口（接口业务流）为例：

1. 【Engine】加载配置文件（系统配置文件）
2. 【Engine】开启系统API监听（默认不开启）
3. 【Engine】加载插件（在当前目录下的mods或plugins目录中搜索）
4. 【Engine】根据配置文件或默认值加载流水线（协程）
5. 【AssemblyLine】检测流水线类型（服务、定时任务）
    - 服务
        - 开启Serve
        - 监听指定的Path
    - 定时任务
        - 开启定时任务
6. 【AssemblyLine】触发流水线
7. 【Worker】执行指定的Steps（Worker - Run方法）
8. 【Worker】根据worker的类型，执行不同的操作（业务、逻辑）
    - 业务
        - 转换参数表达式（如果是服务流）
        - 调用Plugins Run执行具体操作
    - 逻辑
        - 判断
            - 切分三元表达式（判断由三元表达式实现）
            - 将变量表达式转换为具体的值
            - 执行逻辑表达式，获取执行结构
            - 根据执行结构，运行不同的分支
        - 循环
            - 值循环
            - 条件循环
            - 次循环

## 数据结构

### 流水线

蓝图一般是一个独立的文件，其中保存了流水线具体地执行策略，一个蓝图即是一条流水线

蓝图的类型并不固定，支持 json、xml、yml和数据库表类型，只需要符合Nave蓝图规范，即可由Nave转换为统一的蓝图结构体

## 文件类型

- 项目：.pro
- 蓝图：任意，默认为 *.bp


## 表达式

### 参数表达式：主要用请求类流，用于将请求参数转换为执行参数
表达式类似于 `#{name}` ，其中name就是要获取的参数名称
### 变量表达式：获取流中定义的变量
表达式类似于 `@{name}` ，其中name就是要获取的变量名称
### 赋值表达式：获取表达式中的值，可引用变量表达式
表达式类似于 `!{name=123}` ，将name变量赋予值123，必须提前存在该变量
### 条件表达式：执行条件
表达式类似于 `${name==122}?a1:b1` ，其中name==122是条件，a1是条件成立后执行的Worker，b1是条件不成立后执行的Worker

## 版本计划：

### 0.1.0

> 计划完成于：2023年9月1日
> 
> 功能校验版本，仅验证此想法是否可以实现

- ✅支持运行json Exec
- ✅支持基础操作
    - ✅打印字符串
    - ✅记录日志
- ✅基础表达式
    - ✅参数表达式
    - ✅条件表达式
- ✅支持http相关操作
    - ✅监听Http端口(Get)
    - ✅根据不同的路由执行不同的操作
- ✅支持基础逻辑操作
    - ✅条件判断
- ✅引入项目、流水线、执行者的概念
- 🤔支持基础MySQL操作
    - 🤔链接到数据库
    - 🤔执行SQL语句
- ⚠️基础流水线变量
- 🤔支持插件系统（动态 Worker - 热插拔）

### 0.2.0

> 计划完成于：2023年9月15日
> 
> 基础校验版，大概有了最基础的功能

- 💡完成基础Nave Artist的设计与开发
    - 💡基础设计器界面
    - 💡输出Json文件
- 💡底层操作执行改为反射方式
- 💡Http监听完善
    - 💡Post
        - 💡Json
        - 💡Form
        - 💡File
    - 💡Put
    - 💡Delete
    - 💡可以返回数据
- 💡流水线变量增强
- 💡支持更多逻辑操作
    - 💡条件循环
    - 💡次数循环
    - 💡内容循环
- ✅核心引擎改为蓝图驱动

### 0.3.0

> 计划完成于：2023年10月1日

- 支持操作系统操作
- 支持文件系统操作
- 完整支持MySQL
- 支持基础SQLServer
- 支持基础Oracle
- 改为标准EL表达式
- 整体架构整理优化
- 流水线变量增强
- Nave Runner配置文件

### 0.4.0

> 计划完成于：2023年10月15日

- 完善数据类型
  - 字符串
  - 布尔
  - 数组
  - 对象
  - 数字
  - 浮点

### 0.5.0

> 2023年11月1日
> 
> 优化更新，大部分之前的问题应该在此版本修复

- 支持Yml Exec
- 支持Xml Exec
- 项目安全策略升级
- 优化执行速度
- 优化内存使用

### 1.0.0

> 计划完成于：2023年11月15日
> 
> 大版本更新，基础功能均以完善，此版本正式开始开源（MIT 许可证）

- 完整支持SQLServer
- 完整支持Oracle
- 支持数据库Exec
- 初步支持集群化（分布式协调器、分布式配置）
- 完成Nave Watchman的设计与开发
- 业务插件的实现

## 更新规则

> version 1.2.3
> 
> 注意，0.x.x版本所有基础版本也可能存在不兼容的情况

1. 大版本，大版本之间Exec可能存在不兼容情况，需要使用升级器进行升级
2. 基础版本：代表新增了特别的功能或者修复了严重的漏洞
3. 小版本：代表修复了Bug，这种情况下会添加小版本号