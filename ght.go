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
		text := "Замученный дорогой, я выбился из сил\nИ в доме лесника я ночлега попросил\nС улыбкой добродушной старик меня впустил\nИ жестом дружелюбным на ужин пригласил\n(Хэй!)\n\nБудь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!\n\nНа улице темнело, сидел я за столом\nЛесник сидел напротив, болтал о том, о сём\nЧто нет среди животных у старика врагов\nЧто нравится ему подкармливать волков\n\nБудь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!\n\nИ волки среди ночи завыли под окном\nСтарик заулыбался и вдруг покинул дом\nНо вскоре возвратился с ружьём наперевес\n«Друзья хотят покушать, пойдём, приятель, в лес!»\n\nБудь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!"

		songDetail := SongDetail{
			ReleaseDate: "16.07.2006",
			// Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight\n\nI thought I was a fool for no one\nBut, ooh, baby, I'm a fool for you\nYou're the queen of the superficial\nHow long before you tell the truth?\n\nOoh, you set my soul alight\nOoh, you set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
			Text: text,
		}

		c.JSON(http.StatusOK, songDetail)
	})

	r.Run(":8081")
}