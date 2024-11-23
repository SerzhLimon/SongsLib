package transport

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SerzhLimon/SongsLib/internal/models"
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

func (api *Server) SetSong(c *gin.Context) {
	var request models.SetSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}
	if request.Group == "" || request.SongName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'group' and 'song' fields must be provided",
		})
		return
	}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8081/info", nil)
	q := req.URL.Query()
	q.Add("group", request.Group)
	q.Add("song", request.SongName)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from external API: %d", resp.StatusCode)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch valid song info",
		})
		return
	}

	var songInfo models.InfoSong
	if err := json.NewDecoder(resp.Body).Decode(&songInfo); err != nil {
		log.Printf("Error decoding response JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse song info",
		})
		return
	}

	api.Usecase.SetSong(songInfo)

	c.Status(http.StatusCreated)
}

