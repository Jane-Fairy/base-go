## Go依赖管理

go.mod是go语言第三方依赖管理的一种方式，他用来管理当前项目所使用的第三方包，依赖包.



### 1、module

go的module：一组`package`的集合即为一个module，一个module下的`package`版本号相同。一个项目包含一个或者多个module，每个module包含一个或多个 package，每个package包含一个或多个源文件。如图:

![image-20230521054519117](C:\Users\98427\AppData\Roaming\Typora\typora-user-images\image-20230521054519117.png)



一个 module 的版本号规则必须遵循语义化规范：v(major).(minor).(patch)格式。

- **major：**大版本，发生不可兼容的改动时增加该版本
- **minor**：小版本，当有新特性时增加该版本
- **patch**：补丁版本，当有bug修复时增加该版本



用法：

构建module：

```golang
go mod init xxx(module 的名称)

示列：
module github.com/gin-gonic/gin

go 1.13

require (
    github.com/gin-contrib/sse v0.1.0
    github.com/go-playground/validator/v10 v10.9.0
    github.com/goccy/go-json v0.7.8
    github.com/json-iterator/go v1.1.12
    github.com/mattn/go-isatty v0.0.14
    github.com/stretchr/testify v1.7.0
    github.com/ugorji/go/codec v1.2.6
    google.golang.org/protobuf v1.27.1
    gopkg.in/yaml.v2 v2.4.0
)

```

语法：

```
module：声明 module 名称
require：声明依赖及版本号
replace：替换 require 中声明的依赖
exclude：禁用指定依赖
indirect：表示间接依赖
```

