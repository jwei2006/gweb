package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gweb/config"
	"strings"
)

func EntranceFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Conf.IsNeedOpen{
			requestUri := c.Request.RequestURI
			Info().Println(requestUri)
			if !strings.HasSuffix(requestUri, config.Conf.LoginPathWx){
				openid := GetOpenid(c)
				if openid == ""{
					Error().Println("openid空")
					JsonRes(c, config.UnLogin, nil)
					c.Abort()
				}
			}
		}

		c.Next()
	}
}

func UserIdFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.RequestURI, config.Conf.LoginPathWx) {
			c.Next()
			return
		}
		if GetSession(c, config.Key_UserId) == ""{
			JsonRes(c, config.NoEnoughAuth, "u")
			c.Abort()
		}
		c.Next()
	}
}
func CompanyIdFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetSession(c, config.Key_CompanyId) == ""{
			JsonRes(c, config.NoEnoughAuth, "com")
			c.Abort()
		}
		c.Next()
	}
}
