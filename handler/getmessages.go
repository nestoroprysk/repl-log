package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestoroprysk/repl-log/repository"
)

func GetMessages(r *repository.T) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := r.GetMessages()
		c.IndentedJSON(http.StatusOK, result)
	}
}
