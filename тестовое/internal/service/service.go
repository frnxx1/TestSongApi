package service

import (
	"RestSong/internal/database"
	"RestSong/internal/domain/song"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSongInfoQuery(ctx *gin.Context) {
	cursor := ctx.Query("cursor")
	limitStr := ctx.Query("limit")
	limit := 10

	song_title := ctx.Query("name")
	lyrics := ctx.Query("lyrics")
	release_date := ctx.Query("release_date")
	group_name := ctx.Query("group_name")

	query := "SELECT id, group_name , song_title, release_date, lyrics, video_link FROM songs WHERE TRUE"
	args := []interface{}{}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}
	if cursor != "" {
		query += " AND id > $1"
		args = append(args, cursor)
	}

	args = append(args, limit)
	log.Println(args)
	if song_title != "" {
		query += " AND song_title = $3"
		args = append(args, song_title)
	}
	if lyrics != "" {
		query += " AND lyrics = $4"
		args = append(args, lyrics)
	}
	if group_name != "" {
		query += " AND group_name = $5"
		args = append(args, group_name)
	}
	// Использовать отдельный запрос
	if release_date != "" {
		operator := ctx.Query("operator")
		if operator != "" {
			switch operator {
			case "gt":
				query += " AND release_date > $3"
			case "lt":
				query += " AND release_date < $3"
			case "ge":
				query += " AND release_date >= $3"
			case "le":
				query += " AND release_date <= $3"
			}
		}
		args = append(args, release_date)
	}
	query += " ORDER BY id LIMIT $2"
	rows, err := database.GlobalDB.Query(query, args...)
	if err != nil {
		log.Printf("failed get songs with params %+v\n: %v\n", args, err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	defer rows.Close()

	var songs []song.Song
	var nextCursor uint
	for rows.Next() {
		var song song.Song
		if err := rows.Scan(&song.ID, &song.GroupName, &song.SongTitle, &song.ReleaseDate, &song.Lyrics, &song.VideoLink); err != nil {
			log.Printf("failed to scanning songs with params %+v\n: %v\n", args, err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			ctx.Abort()
		}
		songs = append(songs, song)
		nextCursor = song.ID
	}

	ctx.JSON(200, gin.H{
		"songs":       songs,
		"next cursor": nextCursor,
	})
}

func GetVersePagination(ctx *gin.Context) {
	songID := ctx.Param("id")
	amountVerse := ctx.QueryArray("verse")
	if amountVerse == nil {
		log.Printf("Null params ID: %s Params: %v", songID, amountVerse)
		ctx.JSON(404, gin.H{
			"error": "Null query",
		})
	}
	query := "SELECT lyrics FROM songs WHERE id = $1"
	var res []string

	var song song.Song

	err := database.GlobalDB.QueryRow(query, songID).Scan(&song.Lyrics)
	if err != nil {
		log.Printf("failed get songs with ID %s: %v\n", songID, err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
	}

	verseSlice := strings.Split(song.Lyrics, "[Куплет]")
	amountVerse = strings.Split(amountVerse[0], ",")

	sort.Strings(amountVerse)
	if len(amountVerse) > len(verseSlice) {
		log.Println("exceeding the permissible value of couplets")
		ctx.JSON(404, gin.H{"error": "exceeding the permissible value of couplets"})
		ctx.Abort()
	}

	for _, a := range amountVerse {
		str, err := strconv.Atoi(a)
		if err != nil {
			continue
		}
		if str <= 0 {
			continue
		}
		if str > len(verseSlice) {
			continue
		}
		res = append(res, verseSlice[str-1])
	}
	if res == nil {
		ctx.JSON(200, gin.H{
			"message": "Not found verse",
		})
	}
	
	ctx.JSON(200, gin.H{
		"success": res,
	})
}
