 package main

 import (
	 "checkCDN/client/config"
	 "checkCDN/client/control"
	 "fmt"
	 "github.com/kataras/iris"
	 "github.com/kataras/iris/middleware/logger"
	 "github.com/kataras/iris/middleware/recover"
 )

 func router(app  *iris.Application)  {
      apis:=app.Party("/apis")
      {
      	apis.Get("/ping",control.Ping)
	 }

 }
    func main() {
        app := iris.New()
        app.Use(recover.New())
        app.Use(logger.New())
        app.Logger().SetLevel("info")
        router(app)
        listenAddr :=fmt.Sprintf("%s:%d", config.Conf.Get("app.host").(string),config.Conf.Get("app.port").(int64))
        app.Run(iris.Addr(listenAddr), iris.WithoutInterruptHandler)
    }