package main

import (
	"os"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/handler"
	"github.com/nestoroprysk/repl-log/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	c := config.T{Host: os.Getenv("HOST"), Port: os.Getenv("PORT")}

	a, err := client.New(config.T{Host: os.Getenv("SECONDARY_1_HOST"), Port: os.Getenv("SECONDARY_1_PORT")})
	if err != nil {
		panic(err)
	}

	b, err := client.New(config.T{Host: os.Getenv("SECONDARY_2_HOST"), Port: os.Getenv("SECONDARY_2_PORT")})
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

	router.Run(c.Address())
}
