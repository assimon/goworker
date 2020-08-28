
## goworker
goworker(Gin Api 开发骨架)

## 目录结构
```
- bootstrap         // 核心启动
- config            // 常规配置文件
- controller        // 控制器层
- core              // 骨架核心
- middleware        // 中间件
- models            // 模型层
- routers           // 路由层
- runtime           // 缓存及日志文件层
- services          // service层
- static            // 静态资源目录
- storage           // 资源存储
```

## 示例
控制器     
controllers/api/v1/hello.go
```go
package v1
import (
	"github.com/gin-gonic/gin"
	"goworker/core/app"
	"goworker/core/e"
	"net/http"
)

func SayHello(ctx *gin.Context)  {
	gapp := app.Gin{G: ctx}
	gapp.Response(http.StatusOK, e.SUCCESS, "hello, gowrker")
}
```
路由  
routers/routers
```go
package routers

import (
	"github.com/gin-gonic/gin"
	v1 "goworker/controllers/api/v1"
	"goworker/core/config"
	"net/http"
)

func RouterInit() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册图片静态资源目录
	r.StaticFS("/storage/images", http.Dir(config.AppConfig.ImageSavePath))
	r.StaticFS("/storage/files", http.Dir(config.AppConfig.FileSavePath))

	apiv1 := r.Group("api/v1")
	apiv1.GET("/hello", v1.SayHello)

	return r
}
```

## 启动
终端命令
```shell script
$: go run main.go
```

测试请求：
```shell script
http://localhost:8000/api/v1/hello
```

返回：
```json
{"code":200,"msg":"ok","data":"hello, gowrker"}
```