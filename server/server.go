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
		controller.FollowerController
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
		controller.NewFollowerController(s.DB),
	}

	api := s.Echo.Group("/api/v1")

	{
		users := api.Group("/users")
		users.GET("/:id", app.FollowerController.Show)
		users.GET("/new", app.FollowerController.ListNew)
		users.GET("/bye", app.FollowerController.ListBye)
		users.GET("/follow", app.FollowerController.ListFollow)
		users.GET("/followed", app.FollowerController.ListFollower)
		users.GET("/followers/diff/unrequited", app.FollowerController.ListUnrequited)
		//followers.PUT("/init",InitFollowers)
	}
	{
		twitter := api.Group("/twitter")
		twitter.GET("", app.FollowerController.Hoge)
		//tl  := api.Group("/timeline")
		//tl.GET("",Timeline)
	}
}
