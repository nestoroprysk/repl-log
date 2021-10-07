package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"
)

func AppendMessage(r *repository.T, replicas ...*client.T) func(c *gin.Context) {
	return func(c *gin.Context) {
		var m message.T
		if err := json.NewDecoder(c.Request.Body).Decode(&m); err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		if m.Namespace == "" {
			m.Namespace = message.DefaultNamespace
		}

		r.AppendMessage(m)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		errs, ctx := errgroup.WithContext(ctx)
		for _, _rep := range replicas {
			rep := _rep
			errs.Go(func() error {
				return rep.PostMessage(m)
			})
		}

		if err := errs.Wait(); err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusCreated, m)
	}
}
