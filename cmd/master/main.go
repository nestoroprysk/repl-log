package main

import (
	"github.com/nestoroprysk/repl-log/config"

	"github.com/gin-gonic/gin"
)

func main() {
	ping := func(c *gin.Context) { c.Writer.Write([]byte("pong")) }

	router := gin.Default()
	router.GET("/ping", ping)

	c, err := config.Make()
	if err != nil {
		panic(err)
	}

	router.Run(c.Address())
}
