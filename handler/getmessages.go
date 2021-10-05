package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"
)

func GetMessages(r *repository.T) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := r.GetMessages(toNamespaces(c.Request.URL.Query()["namespace"])...)
		c.IndentedJSON(http.StatusOK, result)
	}
}

func toNamespaces(ns []string) []message.Namespace {
	var result []message.Namespace
	for _, n := range ns {
		result = append(result, message.Namespace(n))
	}

	return result
}
