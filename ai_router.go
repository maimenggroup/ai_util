package ai_util

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AIRoute struct {
	address string
	engine  *gin.Engine
	server  *http.Server
}

func (r *AIRoute) Init(address string) error {
	r.address = address
	r.engine = gin.New()
	r.engine.Use(gin.Recovery())
	return nil
}

func (r *AIRoute) SetGetHandler(path string, handlerFunc gin.HandlerFunc) {
	r.engine.GET(path, handlerFunc)
}

func (r *AIRoute) SetPostHandler(path string, handlerFunc gin.HandlerFunc) {
	r.engine.POST(path, handlerFunc)
}

func (r *AIRoute) Run() error {
	// return r.engine.Run(r.address)
	r.server = &http.Server{
		Addr:    r.address,
		Handler: r.engine,
	}
	if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// 优雅的退出
func (r *AIRoute) Close() error {
	// a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.server.Shutdown(ctx)
}
