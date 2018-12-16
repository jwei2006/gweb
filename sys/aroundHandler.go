package sys

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gweb/config"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func Around() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := initCtx(c)
		//initCtx(c)
		e := logEntrance(c, *ctx)
		if e != nil {
			Error().Println(e)
			return
		}

		c.Next()

		/*if c.Writer.Written() {
			return
		}
		c.Request.GetBody*/
	}
}

/*func initCtx(c *gin.Context)  {
	c.Set(config.Key_RequestId, c.Query(config.Key_RequestId))
	c.Set(config.RequestTime, time.Now())
}*/
func initCtx(c *gin.Context) *Context {
	ctx := new(Context)
	ctx.RequestTime = time.Now()
	ctx.PrincipalCode = GetSession(c, config.Key_principal)
	ctx.TenantCode = GetSession(c, config.Key_Tenant)
	ctx.CompanyId = GetSession(c, config.Key_CompanyId)
	ctx.UserId = GetSession(c, config.Key_UserId)
	isManager := GetSession(c, config.Key_Manager)
	if isManager != ""{ctx.IsManager,_ = strconv.Atoi(isManager)}
	ctx.RequestIp = c.ClientIP()
	ctx.RequestId = c.Query(config.Key_RequestId)
	c.Set(config.Key_Ctx, ctx)
	return ctx
}

func logEntrance(c *gin.Context, ctx Context) error {
	var builder strings.Builder
	builder.WriteString("in>")
	builder.WriteString("R:")
	builder.WriteString(ctx.RequestId)
	//builder.WriteString(", P:")
	//builder.WriteString(ctx.PrincipalCode)
	builder.WriteString(", T:")
	builder.WriteString(ctx.TenantCode)
	builder.WriteString(", C:")
	builder.WriteString(ctx.CompanyId)
	builder.WriteString(", U:")
	builder.WriteString(ctx.UserId)
	builder.WriteString(", IP:")
	builder.WriteString(c.ClientIP())
	builder.WriteString(", time:")
	builder.WriteString(strconv.FormatInt(ctx.RequestTime.UnixNano()/1e6, 10))
	builder.WriteString(", M:")
	builder.WriteString(c.Request.Method)
	builder.WriteString(", uri:")
	builder.WriteString(c.Request.RequestURI)
	builder.WriteString(", form:")
	e := c.Request.ParseForm()
	if e == nil {
		Info().Println("!!!!!!!!!!!!!!!!!!!!!!")
		Info().Println(e)
		//return e
	}
	//builder.WriteString(c.Request.PostForm.Encode())
	forms, _ := json.Marshal(c.Request.Form)
	builder.Write(forms)
	builder.WriteString(", jsonBody:")
	jsonBody, _ := ioutil.ReadAll(c.Request.Body)
	builder.WriteString(string(jsonBody))
	builder.WriteString(", params:")
	querys, _ := json.Marshal(c.Request.URL.Query())
	builder.Write(querys)
	Info().Println(builder.String())
	return nil
}
