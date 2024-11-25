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

func (s *Server) SetSong(c *gin.Context) {
	var request models.SetSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}
	if request.Group == "" || request.SongName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'group' and 'song' fields must be provided"})
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

	songInfo := models.InfoSong{
		SongName: request.SongName,
		Group: request.Group,
	}
	if err := json.NewDecoder(resp.Body).Decode(&songInfo); err != nil {
		log.Printf("Error decoding response JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse song info",
		})
		return
	}

	if err = s.Usecase.SetSong(songInfo); err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	
	c.Status(http.StatusCreated)
}

func (s *Server) GetSong(c *gin.Context) {
	var request models.GetSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if request.SongName == "" || request.Offset < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect request"})
		return
	}

	res, err := s.Usecase.GetSong(request)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err,})
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) GetLib(c *gin.Context) {
	var request models.GetLibRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if request.Offset < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect request"})
		return
	}

	res, err := s.Usecase.GetLib(request)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err,})
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteSong(c *gin.Context) {
	var request models.DeleteSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if request.TrackID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'id' fields must be provided"})
		return
	}

	err := s.Usecase.DeleteSong(request)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err,})
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) UpdateSongInfo(c *gin.Context) {
	var request models.UpdateSongInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if request.TrackID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'id' field must be provided"})
		return
	}

	err := s.Usecase.UpdateSongInfo(request)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err,})
	}

	c.Status(http.StatusOK)
}

func (s *Server) UpdateSongText(c *gin.Context) {
	var request models.UpdateSongTextRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if request.TrackID < 1 || request.CoupletNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'id' and 'coupletnum' fields must be provided"})
		return
	}

	err := s.Usecase.UpdateSongText(request)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrcet data"})
	}

	c.Status(http.StatusOK)
}

