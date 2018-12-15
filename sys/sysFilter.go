package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gweb/config"
	"github.com/jwei2006/gweb/util"
	"strings"
)

func entranceFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Conf.IsNeedOpen{
			requestUri := c.Request.RequestURI
			if !strings.HasSuffix(requestUri, config.Conf.LoginPathWx){
				openid := util.GetOpenid(c)
				if openid == ""{
					Debug().Println("openid空")
					JsonRes(c, config.UnLogin, nil)
					c.Abort()
				}
			}
		}

		c.Next()
	}
}

func userIdFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.RequestURI, "/login/wx") {
			c.Next()
			return
		}
		if util.GetSession(c, config.Key_UserId) == ""{
			JsonRes(c, config.NoEnoughAuth, nil)
			c.Abort()
		}
		c.Next()
	}
}
func companyIdFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.GetSession(c, config.Key_CompanyId) == ""{
			JsonRes(c, config.NoEnoughAuth, nil)
			c.Abort()
		}
		c.Next()
	}
}
