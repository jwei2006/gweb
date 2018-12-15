package api

import (
	"github.com/gin-gonic/gin"
	"gweb/util"
	"gweb/config"
	"gweb/service"
	"gweb/sys"
)

func SendPhoneCode(c *gin.Context)  {
	phone := c.PostForm("phone")
	if phone == ""{
		sys.JsonResErr(c, config.ParamError, "手机号缺失", nil)
		return
	}
	code := service.CreateCode()
	util.SetSession(c, "code", code)
	ctx := util.GetCtx(c)
	e := service.SendCode(ctx, phone)
	if e != nil{
		sys.Error().Println(e)
		sys.JsonResErr(c, config.SysError, "发送失败，请稍候重试", nil)
		return
	}
	
}