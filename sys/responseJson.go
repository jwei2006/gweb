package sys

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gweb/config"
	"github.com/jwei2006/gweb/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func JsonRes(c *gin.Context, errorCode int, data interface{}) {
	code := config.GetErrorCode(errorCode)
	JsonResErr(c, code, config.GetErrorTitle(code), data)
}
func JsonResSuccess(c *gin.Context, data interface{}) {
	JsonResErr(c, config.Success, config.GetErrorTitle(config.Success), data)
}
func JsonResErr(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"datas":   data,
		"message": msg,
	})
	ctx := util.GetCtx(c)
	var builder strings.Builder
	builder.WriteString("out<")
	builder.WriteString("R:")
	builder.WriteString(ctx.RequestId)
	//builder.WriteString(", P:")
	//builder.WriteString(ctx.PrincipalCode)
	builder.WriteString(", T:")
	builder.WriteString(ctx.TenantCode)
	builder.WriteString(", U:")
	builder.WriteString(ctx.UserId)
	builder.WriteString(", time:")
	builder.WriteString(strconv.FormatInt(ctx.RequestTime.UnixNano()/1e6, 10))
	builder.WriteString(", resTime:")
	builder.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
	builder.WriteString(", uri:")
	builder.WriteString(c.Request.RequestURI)
	builder.WriteString(", code:")
	builder.WriteString(strconv.Itoa(code))
	builder.WriteString(", msg:")
	builder.WriteString(msg)
	builder.WriteString(", data:")
	bytes, _ := json.Marshal(data)
	builder.WriteString(string(bytes))
	Info().Println(builder.String())
}
