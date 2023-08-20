package routes

import (
	"chat_real_time_go/src/api/v1/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	//db  *sql.DB
}

func NewRouter(eng *gin.Engine) Router {
	return &router{eng: eng}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.rg.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.websockets()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) websockets() {

	r.rg.GET("/ws", func(c *gin.Context) {
		handler.HandleConnections(c)
	})
	r.eng.StaticFile("/websockets.html", "./assets/static/websockets.html")
}
