package main

import (
	"fmt"
	"log"

	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/handler"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"
	"github.com/nestoroprysk/repl-log/util"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	replicas, err := util.ToClients(c.Replicate)
	if err != nil {
		log.Fatal(err)
	}

	r := repository.New()

	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.GET("/messages", handler.GetMessages(r))
	router.POST("/messages", handler.AppendMessage(r, replicas...))
	router.GET("/namespaces", handler.GetNamespaces(r))
	router.DELETE(fmt.Sprintf("/namespaces/:%s", message.NamespaceID), handler.DeleteNamespace(r, replicas...))

	router.Run(c.Listen.Address())
}
