package util

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gweb/config"
	"gweb/sys"
)

func GetSession(c *gin.Context, key string) string {
	session := sessions.Default(c)
	val := session.Get(key)
	if val == nil{return ""}
	return val.(string)
}

func SetSession(c *gin.Context, key, val string)  {
	session := sessions.Default(c)
	session.Set(key, val)
	session.Save()
}

func GetCtx(c *gin.Context) sys.Context {
	ctx, exists := c.Get(config.Key_Ctx)
	if !exists {
		sys.Error().Fatalln(errors.New("上下文环境不存在"))
		_ctx := new(sys.Context)
		_ctx.UserId = ""
		return *_ctx
	}
	v, ok := ctx.(sys.Context)
	if ok{
		return v
	}
	_ctx := new(sys.Context)
	_ctx.UserId = ""
	return *_ctx
}

func GetOpenid(c *gin.Context) string {
	return GetSession(c, config.Openid)
}