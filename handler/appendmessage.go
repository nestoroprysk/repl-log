package handler

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/repository"
)

type Option func(message.T) error

func AppendMessage(r *repository.T, opts ...Option) func(c *gin.Context) {
	return func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		m := message.T(b)
		r.AppendMessage(m)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		errs, ctx := errgroup.WithContext(ctx)
		for _, opt := range opts {
			o := opt
			errs.Go(func() error {
				return o(m)
			})
		}

		if err := errs.Wait(); err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusCreated, m)
	}
}

func Replicate(c *client.T) Option {
	return func(m message.T) error {
		return c.PostMessage(m)
	}
}