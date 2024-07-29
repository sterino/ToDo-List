package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"todo-list/internal/api/handler"
	"todo-list/internal/api/routes"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(taskHandler *handler.TaskHandler) *Server {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MethodNotAllowedMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.InitRoutes(router.Group("/api"), taskHandler)

	return &Server{router}
}

func (s *Server) Run(infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Printf("starting server on: 8080")
	err := s.engine.Run(":8080")
	errorLog.Fatal(err)
}

func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedMethods := map[string]bool{
			"GET":    true,
			"POST":   true,
			"PUT":    true,
			"DELETE": true,
		}

		if !allowedMethods[c.Request.Method] {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "Method Not Allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
