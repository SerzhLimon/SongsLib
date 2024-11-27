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

const (
	Peremen = "Вместо тепла - зелень стекла\nВместо огня - дым\nИз сетки календаря выхвачен день\nКрасное солнце сгорает дотла\nДень догорает с ним\nНа пылающий город падает тень\n\nПеремен требуют наши сердца\nПеремен требуют наши глаза\nВ нашем смехе и в наших слезах, и в пульсации вен\nПеремен, мы ждём перемен\n\nЭлектрический свет продолжает наш день\nИ коробка спичек пуста\nНо на кухне синим цветком горит газ\nСигареты в руках, чай на столе, эта схема проста\nИ больше нет ничего, всё находится в нас\n\nПеремен требуют наши сердца\nПеремен требуют наши глаза\nВ нашем смехе и в наших слезах, и в пульсации вен\nПеремен, мы ждём перемен\n\nМы не можем похвастаться мудростью глаз\nИ умелыми жестами рук\nНам не нужно всё это, чтобы друг друга понять\nСигареты в руках, чай на столе, так замыкается круг\nИ вдруг нам становится страшно что-то менять\n\nПеремен требуют наши сердца\nПеремен требуют наши глаза\nВ нашем смехе и в наших слезах, и в пульсации вен\nПеремен, мы ждём перемен"
	Zvezda = "Белый снег, серый лед, на растрескавшейся земле.\nОдеялом лоскутным на ней - город в дорожной петле.\nА над городом плывут облака, закрывая небесный свет.\nА над городом - желтый дым, городу две тысячи лет,\nПрожитых под светом Звезды по имени Солнце...\n\nИ две тысячи лет - война, война без особых причин.\nВойна - дело молодых, лекарство против морщин.\nКрасная, красная кровь - через час уже просто земля,\nЧерез два на ней цветы и трава, через три она снова жива\nИ согрета лучами Звезды по имени Солнце...\n\nИ мы знаем, что так было всегда, что Судьбою больше любим,\nКто живет по законам другим и кому умирать молодым.\nОн не помнит слово \"да\" и слово \"нет\", он не помнит ни чинов, ни имен.\nИ способен дотянуться до звезд, не считая, что это сон,\nИ упасть, опаленным Звездой по имени Солнце..."

	Lesnik  = "Замученный дорогой, я выбился из сил\nИ в доме лесника я ночлега попросил\nС улыбкой добродушной старик меня впустил\nИ жестом дружелюбным на ужин пригласил\n(Хэй!)\n\nБудь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!\n\nНа улице темнело, сидел я за столом\nЛесник сидел напротив, болтал о том, о сём\nЧто нет среди животных у старика врагов\nЧто нравится ему подкармливать волков\n\nБудь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!\n\nИ волки среди ночи завыли под окном\nСтарик заулыбался и вдруг покинул дом\nНо вскоре возвратился с ружьём наперевес\n«Друзья хотят покушать, пойдём, приятель, в лес!»\n\nБудь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!"
	DvaVora = "Два вора, лихо скрывшись от погони\nДелить украденное золото решили\nНа старом кладбище, вечернею порою\nУселись рядом на заброшенной могиле\nИ вроде поровну досталось им богатство\nНо вот беда — последняя монета\nОдин кричит: «Она моя — я лучше дрался!»\n«Да что б ты делал, друг, без моего совета?»\n\n— Отдай монету, а не то я рассержусь\n— Мне наплевать, я твоей злости не боюсь\n— Но ведь я похитил деньги и всё дело провернул\n— Без моих идей, невежа, ты бы и шагу не шагнул\n\nЧто же делать нам с монетой, как же нам её делить?\n— Отдадим покойнику\n— Отлично! Так тому и быть\n\n— Я был проворней, значит, денежка моя\n— Не допущу, чтоб ты богаче был, чем я\n— Сейчас вцеплюсь тебе я в горло и на части разорву\n— Я прибью тебя дубиной и все деньги заберу!\n\nЧто же делать нам с монетой, как же нам её делить?\n— Отдадим покойнику\n— Отлично! Так тому и быть\n\nИ мертвец, гремя костями, вдруг поднялся из земли:\n«Довели меня, проклятые, ей-богу, довели!\nВоры вмиг переглянулись, и помчались наутёк\nА мертвец всё золото с собой в могилу уволок\n"
	Dom     = "В заросшем парке стоит старинный дом\nЗабиты окна, и мрак царит извечно в нём\nСказать я пытался: «Чудовищ нет на земле»\nНо тут же раздался ужасный голос во мгле\nГолос во мгле\n\n«Мне больно видеть белый свет, мне лучше в полной темноте\nЯ очень много-много лет мечтаю только о еде\nМне слишком тесно взаперти, и я мечтаю об одном:\nСкорей свободу обрести, прогрызть свой ветхий старый дом\nПроклятый старый дом»\n\nБыл дед, да помер, слепой и жутко злой\nНикто не вспомнил о нём с зимы холодной той\nСоседи не стали его тогда хоронить\nЛишь доски достали, решили заколотить\nДвери и окна\n\n«Мне больно видеть белый свет, мне лучше в полной темноте\nЯ очень много-много лет мечтаю только о еде\nМне слишком тесно взаперти, и я мечтаю об одном:\nСкорей свободу обрести, прогрызть свой ветхий старый дом\nПроклятый старый дом»\n\nИ это место стороной обходит сельский люд\nИ суеверные твердят: «Там призраки живут»\n"
)

func bind() map[string][3]string {
	res := make(map[string][3]string)

	res["Лесник"] = [3]string{Lesnik, "1.02.2003", "https://genius.com/Korol-i-shut-forester-lyrics"}
	res["Проклятый старый дом"] = [3]string{Dom, "1.03.2004", "https://genius.com/Korol-i-shut-the-cursed-old-house-lyrics"}
	res["Два вора и монета"] = [3]string{DvaVora, "1.04.2005", "https://genius.com/Korol-i-shut-two-thieves-and-a-coin-lyrics"}

	res["Звезда по имени солнце"] = [3]string{Zvezda, "5.10.1980", "https://genius.com/Kino-star-called-sun-lyrics"}
	res["Перемен"] = [3]string{Zvezda, "18.12.1981", "https://genius.com/Kino-changes-lyrics"}

	return res
}

func main() {
	r := gin.Default()

	songs := bind()

	r.GET("/info", func(c *gin.Context) {
		group := c.Query("group")
		song := c.Query("song")

		if group == "" || song == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Group and song parameters are required"})
			return
		}

		if _, ok := songs[song]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unknown track"})
			return
		}
		songDetail := SongDetail{
			ReleaseDate: songs[song][1],
			Link: songs[song][2],
			Text: songs[song][0],
		}

		c.JSON(http.StatusOK, songDetail)
	})

	r.Run(":8081")
}