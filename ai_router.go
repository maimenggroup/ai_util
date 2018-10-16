package ai_util

import (
	"github.com/gin-gonic/gin"
)

type AIRoute struct {
	address string
	engine  *gin.Engine
}

func (r *AIRoute) Init(address string) error {
	r.address = address
	r.engine = gin.Default()
	return nil
}

func (r *AIRoute) SetGetHandler(path string, handlerFunc gin.HandlerFunc)  {
	r.engine.GET(path, handlerFunc)
}

func (r *AIRoute) SetPostHandler(path string, handlerFunc gin.HandlerFunc)  {
	r.engine.POST(path, handlerFunc)
}

func (r *AIRoute) Run() error {
	return r.engine.Run(r.address)
}
