package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/adaptor/controller"
	"github.com/k-ueki/tmanager/server/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	App struct {
		controller.FollowerQueryController
		controller.FollowerCommandController
	}

	Server struct {
		DB   *gorm.DB
		Echo *echo.Echo
	}
)

const (
	clientPort = ":6060"
	port       = ":7777"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func NewServer() (*Server, error) {
	db, err := util.NewDB()
	if err != nil {
		return nil, err
	}
	return &Server{
		DB:   db,
		Echo: echo.New(),
	}, nil
}

func run() error {
	server, err := NewServer()
	if err != nil {
		return err
	}

	server.Echo.Use(middleware.CORS())
	//server.Echo.HTTPErrorHandler()

	setRoutes(server)

	return server.Echo.Start(port)
}

func setRoutes(s *Server) {
	app := &App{
		controller.NewFollowerQueryController(s.DB),
		controller.NewFollowerCommandController(s.DB),
	}

	api := s.Echo.Group("/api/v1")
	{
		users := api.Group("/users")
		users.POST("/init", app.FollowerCommandController.Init)
		users.GET("/:id", app.FollowerQueryController.Show)
		users.GET("/new", app.FollowerQueryController.ListNew)
		users.GET("/bye", app.FollowerQueryController.ListBye)
		users.GET("/follow", app.FollowerQueryController.ListFollow)
		users.GET("/followed", app.FollowerQueryController.ListFollower)
		users.GET("/followers/diff/unrequited", app.FollowerQueryController.ListUnrequited)
		//followers.PUT("/init",InitFollowers)
	}
	{
		twitter := api.Group("/twitter")
		twitter.GET("", app.FollowerQueryController.Hoge)
		//tl  := api.Group("/timeline")
		//tl.GET("",Timeline)
	}
}
