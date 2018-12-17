package sys

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jwei2006/gweb/config"
	"io"
	"os"
	"strconv"
)

func Run(initRouter func(*gin.Engine), findPrincipal func(string) string) {
	InitLogger()
	gin.DisableConsoleColor()
	if !config.Conf.IsDebug {
		file, _ := os.Create(config.Conf.LogGinPath)
		gin.DefaultWriter = io.MultiWriter(file)
	}
	InitDb()
	defer Close()
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(config.Conf.CookieId))
	router.Use(sessions.Sessions(config.Conf.SessionId, store))
	router.Use(Around(findPrincipal))
	initRouter(router)

	router.Run(":" + strconv.Itoa(config.Conf.Port))
}
