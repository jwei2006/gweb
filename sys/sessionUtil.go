package sys

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gweb/config"
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

func GetCtx(c *gin.Context) Context {
	ctx, exists := c.Get(config.Key_Ctx)
	if !exists {
		Error().Fatalln(errors.New("上下文环境不存在"))
		_ctx := new(Context)
		return *_ctx
	}
	if v, ok := ctx.(*Context); ok{
		return *v
	}
	Error().Fatalln(errors.New("获取上下文环境失败"))
	_ctx := new(Context)
	return *_ctx
}

func GetOpenid(c *gin.Context) string {
	return GetSession(c, config.Openid)
}