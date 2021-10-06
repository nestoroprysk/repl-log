package main

import (
	"fmt"

	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/handler"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := repository.New()

	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.GET("/messages", handler.GetMessages(r))
	router.POST("/messages", handler.AppendMessage(r))
	router.GET("/namespaces", handler.GetNamespaces(r))
	router.DELETE(fmt.Sprintf("/namespaces/:%s", message.NamespaceID), handler.DeleteNamespace(r))

	router.Run(config.Secondary.Address())
}
