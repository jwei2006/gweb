package sys

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"gweb/config"
	"strconv"
)

func Run(initRouter func(*gin.Engine))  {
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
	router.Use(Around())
	initRouter(router)

	router.Run(":" + strconv.Itoa(config.Conf.Port))
}
