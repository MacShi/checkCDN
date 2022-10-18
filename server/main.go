 package main

 import (
	 "checkCDN/server/config"
	 "checkCDN/server/control"
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
    	//aa :=config.Conf.Get("client.urls").([]interface{})
    	//fmt.Println(aa[0])
    	//urls := config.Conf.Get("client.infos").([]*toml.Tree)
    	//for _,value :=range urls{
    	//	fmt.Println(value.Get("host").(string))
		//}



    }