package control

import (
	"checkCDN/server/config"
	"checkCDN/server/service"
	"checkCDN/server/utils"
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