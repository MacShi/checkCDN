package control

import (
	"checkCDN/client/config"
	"checkCDN/client/service"
	"checkCDN/client/utils"
	"github.com/kataras/iris"
)
func Ping(ctx iris.Context)  {
	token :=config.Conf.Get("app.token").(string)
	if token == ctx.URLParam("token"){
		domain := ctx.URLParam("host")
		result := service.Ping(domain)
		ctx.JSON(utils.ResultUtil.Success(result))
	}else {
		ctx.JSON(utils.ResultUtil.Failure("认证失败"))
	}
}