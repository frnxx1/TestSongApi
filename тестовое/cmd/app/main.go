package main

import (
	"RestSong/internal/database"
	"RestSong/internal/domain/song"
	"RestSong/internal/service"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "RestSong/cmd/app/docs"
)


// @title Song API
// @version 1.0
// @description This is a sample server for a song API
// @host localhost:3001
// @BasePath /crud
func handlerFunc() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Для удобства, что бы не запутаться. А вообще мог бы разбить по файлам, но запросов мало, так что не стал
	api := r.Group("/crud")
	{
		api.POST("/song", service.AddSongInfo)
		api.GET("/songs", service.GetSongsInfo)
		api.PUT("/song/:id", service.UpdateSongInfo)
		api.DELETE("/song/:id", service.DeleteSongInfo)
	}
	// Запросы с параметрами
	r.GET("/song", service.GetSongInfoQuery)
	r.GET("/song/:id", service.GetVersePagination)
	// Для заполнения таблиц
	r.GET("/songs/seed", func(c *gin.Context) {
		var song song.Song

		for i := 1; i <= 10; i++ {
			song.SongTitle = fmt.Sprintf("Song%d", i)
			song.Lyrics = fmt.Sprintf("text%d", i)
			song.GroupName = fmt.Sprintf("authors%d", i)
			song.VideoLink = fmt.Sprintf("Link%d", i)
			song.ReleaseDate = time.Now().Add(-time.Duration(22-i) * time.Hour)

			_, err := database.GlobalDB.Exec("INSERT INTO songs (song_title, group_name, lyrics, video_link, release_date) VALUES ($1 ,$2, $3, $4, $5)", song.SongTitle, song.GroupName, song.Lyrics, song.VideoLink, song.ReleaseDate)
			time.Sleep(time.Second)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		c.JSON(200, gin.H{
			"message": song,
		})
	})

	return r
}

func main() {
	// Загрузка .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed load env, err: ", err.Error())
	}
	db := database.InitDatabase()
	defer db.Close()

	router := handlerFunc()
	router.Run(":3001")
}
