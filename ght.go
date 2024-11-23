package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func main() {
	r := gin.Default()

	r.GET("/info", func(c *gin.Context) {
		group := c.Query("group")
		song := c.Query("song")

		if group == "" || song == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Group and song parameters are required"})
			return
		}
		// text := "I thought I was a fool for no one\nBut, ooh, baby, I'm a fool for you\nYou're the queen of the superficial\nHow long before you tell the truth?\n\nOoh, you set my soul alight\nOoh, you set my soul alight"
		// Для примера, возвращаем фиксированные данные
		songDetail := SongDetail{
			ReleaseDate: "16.07.2006",
			Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight\n\nI thought I was a fool for no one\nBut, ooh, baby, I'm a fool for you\nYou're the queen of the superficial\nHow long before you tell the truth?\n\nOoh, you set my soul alight\nOoh, you set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		}

		c.JSON(http.StatusOK, songDetail)
	})

	r.Run(":8081")
}