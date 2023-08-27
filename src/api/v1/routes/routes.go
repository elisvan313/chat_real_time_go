package routes

import (
	"chat_real_time_go/internal/users"
	"chat_real_time_go/src/api/v1/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *gorm.DB
}

func NewRouter(eng *gin.Engine, db *gorm.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.rg.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.websockets()
	r.users()
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
func (r *router) users() {
	group := r.rg.Group("/users")
	repo := users.NewRepository(r.db)
	service := users.NewService(repo)
	handler := handler.NewUser(service)
	group.GET("/users", handler.GetAll())
	group.POST("/create", handler.CreateUser())
}
