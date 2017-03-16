package main

import (
	"context"
	"fmt"
	"github.com/g8os/go-client"
	"github.com/pmylund/go-cache"
	"net/http"
	"time"
)

const (
	connectionPoolMiddlewareKey = "github.com/g8os/grid+connection-pool"
)

type connectionMiddleware struct {
	handler http.Handler
	pools   *cache.Cache
}

func (c *connectionMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), connectionPoolMiddlewareKey, c)
	r = r.WithContext(ctx)

	c.handler.ServeHTTP(w, r)
}

func ConnectionMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		p := &connectionMiddleware{
			pools: cache.New(5*time.Minute, 1*time.Minute),
		}

		return p
	}
}

func GetConnection(r *http.Request, id string) client.Client {
	return client.NewClient(fmt.Sprintf("%s:6379", id), "")
}
