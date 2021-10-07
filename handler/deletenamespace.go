package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"
)

func DeleteNamespace(r *repository.T, replicas ...*client.T) func(c *gin.Context) {
	return func(c *gin.Context) {
		n := message.Namespace(c.Params.ByName(message.NamespaceID))

		if ok := r.DeleteNamespace(n); ok == false {
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		errs, ctx := errgroup.WithContext(ctx)
		for _, _rep := range replicas {
			rep := _rep
			errs.Go(func() error {
				_, err := rep.DeleteNamespace(n)
				return err
			})
		}

		if err := errs.Wait(); err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
	}
}
