## на случай, если что-то пойдет не так

- файле internal/transport/transport.go раскомментируйте 58 строку, запустите файл fortest/fortest.go
- запустите make
- дернув эндоинт /setsong можно заполнить библиотеку 5 готовыми треками:

curl -X POST "http://localhost:8080/setsong" \
-H "Content-Type: application/json" \
-d '{
  "group": "Король и шут",
  "song": "Два вора и монета"      
}'

curl -X POST "http://localhost:8080/setsong" \
-H "Content-Type: application/json" \
-d '{
  "group": "Король и шут",
  "song": "Лесник"      
}'

curl -X POST "http://localhost:8080/setsong" \
-H "Content-Type: application/json" \
-d '{
  "group": "Король и шут",
  "song": "Проклятый старый дом"      
}'

curl -X POST "http://localhost:8080/setsong" \
-H "Content-Type: application/json" \
-d '{
  "group": "Кино",
  "song": "Перемен"      
}'

curl -X POST "http://localhost:8080/setsong" \
-H "Content-Type: application/json" \
-d '{
  "group": "Кино",
  "song": "Звезда по имени солнце"      
}'

## /getlib выводит основную информацию о песнях, без текста 

curl -X GET "http://localhost:8080/getlib" \
-H "Content-Type: application/json" \
-d '{
  "offset": 2
}'

{
  "Songs": [
    {
      "ID": 18,
      "SongName": "Перемен",
      "Group": "Кино",
      "ReleaseDate": "18.12.1981",
      "Link": "https://genius.com/Kino-changes-lyrics"
    },
    {
      "ID": 19,
      "SongName": "Звезда по имени солнце",
      "Group": "Кино",
      "ReleaseDate": "5.10.1980",
      "Link": "https://genius.com/Kino-star-called-sun-lyrics"
    }
  ]
}

## /getsong выводит текст песни с пагинацией по 1 куплету

curl -X GET "http://localhost:8080/getsong" \
-H "Content-Type: application/json" \
-d '{
  "songname": "Лесник",
  "offset": 2
}'

{
  "CoupletNumber": 2,
  "Couplet": "Будь как дома, путник\nЯ ни в чём не откажу\nЯ ни в чём не откажу\nЯ ни в чём не откажу! (Хэй!)\nМножество историй\nКоль желаешь, расскажу\nКоль желаешь, расскажу\nКоль желаешь, расскажу!"
}

## /updatesonginfo обновляет основную информацию, без текста песни. Все поля опциональны, кроме id 

curl -X PATCH "http://localhost:8080/updatesonginfo" \
-H "Content-Type: application/json" \
-d '{
  "id": 9,
  "songname": "!!!!Проклятый старый дом!!!!",
  "link": "empty link",
  "releasedate": "07.07.2077"
}'

## /updatesongtext обновляет куплет песни. При обновлении указываем номер куплета для изменения 

curl -X PATCH "http://localhost:8080/updatesongtext" \
-H "Content-Type: application/json" \
-d '{
  "id": 8,
  "coupletnum": 1,
  "text": "Измененный куплет"
}'
