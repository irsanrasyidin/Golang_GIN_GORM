package handler

import (
	"Golang_GIN_GORM/config"
	"Golang_GIN_GORM/manager"
	"Golang_GIN_GORM/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv  *gin.Engine
	host string
}

func (s *server) Run() {
	// session
	// store := cookie.NewStore([]byte("secret"))

	// s.srv.Use(sessions.Sessions("session", store))

	s.srv.Use(middleware.LoggerMiddleware())

	// handler
	NewLogicHandler(s.srv, s.usecaseManager.GetLogicUsecase())
	NewPostgreSqlHandler(s.srv, s.usecaseManager.GetPostgreSqlUsecase())
	NewMainHandler(s.srv, s.usecaseManager.GetMainUsecase())

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.NewConfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()
	srv.Static("/assets", "scr/assets/")
	srv.LoadHTMLGlob("scr/html/*")

	if c.DbConfig.Host == "" || c.AppPort == "" {
		panic("No Host or port define")
	}

	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
