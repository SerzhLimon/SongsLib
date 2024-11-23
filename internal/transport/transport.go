package transport

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SerzhLimon/SongsLib/internal/repository"
	uc "github.com/SerzhLimon/SongsLib/internal/usecase"
)



type Server struct {
	Usecase uc.UseCase
}

func NewServer(database *sql.DB) *Server {
	pgClient := repository.NewPGRepository(database)
	uc := uc.NewUsecase(pgClient)

	return &Server{
		Usecase: uc,
	}
}

func (api *Server) GetSongs(c *gin.Context) {

	c.Status(http.StatusOK)
}
