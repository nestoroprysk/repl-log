package main

import (
	"fmt"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/handler"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	a, err := client.New(config.SecondaryA)
	if err != nil {
		panic(err)
	}

	b, err := client.New(config.SecondaryB)
	if err != nil {
		panic(err)
	}

	r := repository.New()

	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.GET("/messages", handler.GetMessages(r))
	router.POST("/messages", handler.AppendMessage(r,
		handler.Replicate(a),
		handler.Replicate(b),
	))
	router.GET("/namespaces", handler.GetNamespaces(r))
	router.DELETE(fmt.Sprintf("/namespaces/:%s", message.NamespaceID), handler.DeleteNamespace(r,
		handler.ReplicateNamespace(a),
		handler.ReplicateNamespace(b),
	))

	router.Run(config.Master.Address())
}
