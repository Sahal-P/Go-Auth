package api

import (
	"net/http"

	"github.com/Sahal-P/Go-Auth/db"
	"github.com/Sahal-P/Go-Auth/service/user"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *db.PostgreSQLStorage
}

func NewAPIServer(addr string, db *db.PostgreSQLStorage) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(apiV1)

	return http.ListenAndServe(s.addr, router)
}
