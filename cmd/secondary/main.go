package main

import (
	"os"

	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/handler"
	"github.com/nestoroprysk/repl-log/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := repository.New()

	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.GET("/messages", handler.GetMessages(r))
	router.POST("/messages", handler.AppendMessage(r))

	c := config.T{Host: os.Getenv("HOST"), Port: os.Getenv("PORT")}
	router.Run(c.Address())
}
