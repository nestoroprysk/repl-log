package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestoroprysk/repl-log/repository"
)

func GetNamespaces(r *repository.T) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := r.GetNamespaces()
		c.IndentedJSON(http.StatusOK, result)
	}
}
