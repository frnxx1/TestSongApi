package service

import (
	"RestSong/internal/database"
	"RestSong/internal/domain/song"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSongsInfo(ctx *gin.Context) {
	query := "SELECT id, group_name , song_title, release_date, lyrics, video_link FROM songs WHERE TRUE"
	rows, err := database.GlobalDB.Query(query)
	if err != nil {
		log.Printf("failed get songs %v\n", err)
		ctx.JSON(500, gin.H{"error": err.Error})
		ctx.Abort()
	}

	var songs []song.Song

	for rows.Next() {
		var song song.Song
		if err := rows.Scan(&song.ID, &song.GroupName, &song.SongTitle, &song.ReleaseDate, &song.Lyrics, &song.VideoLink); err != nil {
			log.Printf("failed scanning songs %v\n", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			ctx.Abort()
		}
		songs = append(songs, song)
	}

	ctx.JSON(200, gin.H{
		"songs": songs,
	})
}

// AddSongInfo godoc
// @Summary Add a new song
// @Description Add a new song to the database
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body song.Song true "Add song"
// @Success 201 {object} song.Song
// @Failure 400 {object} gin.H
// @Router /crud/song [post]
func AddSongInfo(ctx *gin.Context) {
	var song song.Song
	if err := ctx.ShouldBindJSON(&song); err != nil {
		log.Printf("failed get song data %v\n", err)
		ctx.JSON(404, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	song.ReleaseDate = time.Now()
	query := "INSERT INTO songs (group_name , song_title, release_date, lyrics, video_link) VALUES ($1, $2, $3, $4, $5)"
	_, err := database.GlobalDB.Exec(query, song.ID, song.GroupName, song.SongTitle, song.ReleaseDate, song.Lyrics, song.VideoLink)
	if err != nil {
		log.Printf("failed add song with data %+v\n: %v\n", song, err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	ctx.JSON(200, gin.H{
		"success": song,
	})
}

func UpdateSongInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	var song song.Song
	if err := ctx.ShouldBindJSON(&song); err != nil {
		log.Printf("failed get song data %v\n", err)
		ctx.JSON(404, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	query := "UPDATE songs SET group_name = $1, song_title = $2, release_date = $3, lyrics = $4, video_link = $5 WHERE id = $6"

	_, err := database.GlobalDB.Exec(query, song.GroupName, song.SongTitle, song.ReleaseDate, song.Lyrics, song.VideoLink, id)
	if err != nil {
		log.Printf("failed update song with ID %s: %v\n", id, err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	ctx.JSON(200, gin.H{"success": "Success update"})
}

func DeleteSongInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	query := "DELETE FROM songs WHERE id = $1"

	_, err := database.GlobalDB.Exec(query, id)
	if err != nil {
		log.Printf("failed delete song with ID %s: %v\n", id, err)
		ctx.JSON(500, gin.H{"error": err.Error})
		ctx.Abort()
	}

	ctx.JSON(200, gin.H{"success": "Success delete"})
}
