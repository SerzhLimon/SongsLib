package transport

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

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
		logrus.WithError(err).Error("error decode JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if request.Group == "" || request.SongName == "" {
		logrus.Warn("validation failed: 'group' or 'songname' is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'group' and 'song' fields must be provided"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %s %s", request.SongName, request.Group)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8081/info", nil)
	q := req.URL.Query()
	q.Add("group", request.Group)
	q.Add("song", request.SongName)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.WithError(err).Errorf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to set song"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.WithError(err).Errorf("unexpected status code from external API: %d", resp.StatusCode)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch valid song info"})
		return
	}

	songInfo := models.InfoSong{
		SongName: request.SongName,
		Group:    request.Group,
	}
	if err := json.NewDecoder(resp.Body).Decode(&songInfo); err != nil {
		logrus.WithError(err).Error("error decode JSON")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse song info"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %s %s %s %s %s",
		songInfo.SongName, songInfo.Group, songInfo.ReleaseDate, songInfo.Link, songInfo.Text)

	if err = s.Usecase.SetSong(songInfo); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to set song"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "true"})
}

func (s *Server) GetSong(c *gin.Context) {
	var request models.GetSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if request.SongName == "" || request.Offset < 1 {
		logrus.Warn("validation failed: 'offset' must be greater than 0, songname must be not empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect request"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %d %s", request.Offset, request.SongName)

	res, err := s.Usecase.GetSong(request)
	if err != nil {
		logrus.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "unknown song or couplet"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to get song"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) GetLib(c *gin.Context) {
	var request models.GetLibRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if request.Offset < 1 {
		logrus.Warn("validation failed: 'offset' must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect offset"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %d %s %s %s %s",
		request.Offset,
		uc.SafeDereference(&request.SongName),
		uc.SafeDereference(&request.Group),
		uc.SafeDereference(&request.ReleaseDate),
		uc.SafeDereference(&request.Link),
	)

	res, err := s.Usecase.GetLib(request)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to get songs"})
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteSong(c *gin.Context) {
	var request models.DeleteSongRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if request.TrackID < 1 {
		logrus.Warn("validation failed: 'TrackID' must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'id' fields must be provided"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %d", request.TrackID)

	err := s.Usecase.DeleteSong(request)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to delete song"})
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func (s *Server) UpdateSongInfo(c *gin.Context) {
	var request models.UpdateSongInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	if request.TrackID < 1 {
		logrus.Warn("validation failed: 'TrackID' must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{"error": "both 'id' field must be provided"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %d %s %s %s %s",
		request.TrackID,
		uc.SafeDereference(request.NewSongName),
		uc.SafeDereference(request.NewGroup),
		uc.SafeDereference(request.NewReleaseDate),
		uc.SafeDereference(request.NewLink),
	)

	err := s.Usecase.UpdateSongInfo(request)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to update song"})
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func (s *Server) UpdateSongText(c *gin.Context) {
	var request models.UpdateSongTextRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if request.TrackID < 1 || request.CoupletNum < 1 {
		logrus.Warn("validation failed: 'TrackID' and 'CoupletNum' must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both 'id' and 'coupletnum' fields must be provided"})
		return
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("Parsed request: %d %d %s", request.TrackID, request.CoupletNum, uc.SafeDereference(request.NewText))

	err := s.Usecase.UpdateSongText(request)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to update song"})
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
