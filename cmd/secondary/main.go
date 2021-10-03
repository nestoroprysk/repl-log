package main

import (
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

	c, err := config.Make()
	if err != nil {
		panic(err)
	}

	router.Run(c.Address())
}
