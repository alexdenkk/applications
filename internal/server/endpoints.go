package server

import (
	"alexdenkk/applications/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegisterEndpoints() {
	s.Engine.LoadHTMLGlob("web/html/*.html")

	r := s.Engine.Group("/")
	r.GET("/", s.HomePage)
	r.POST("/", s.CreateApplication)
}

func (s *Server) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (s *Server) CreateApplication(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	text := c.PostForm("text")

	if len(name) < 2 || len(phone) < 6 || len(text) == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "too short field"},
		)
	}

	application := &entity.Application{
		Name: name,
		Phone: phone,
		Text: text,
	}

	s.Writer.SendMessage(
		c.Request.Context(),
		application,
	)
}
