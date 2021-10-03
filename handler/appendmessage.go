package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"
)

func AppendMessage(r *repository.T) func(c *gin.Context) {
	return func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		m := message.T(b)
		r.AppendMessage(m)
		c.IndentedJSON(http.StatusCreated, m)
	}
}
