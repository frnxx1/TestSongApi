package song

import "time"

type Song struct {
	ID          uint      `json:"id" db:"id"`                     // Айди
	GroupName   string    `json:"group_name" db:"group_name"`     // Название группы
	SongTitle   string    `json:"song_title" db:"song_title"`     // Название песни
	ReleaseDate time.Time `json:"release_date" db:"release_date"` // Дата выпуска
	Lyrics      string    `json:"lyrics" db:"lyrics"`             // Текст песни
	VideoLink   string    `json:"video_link" db:"video_link"`     // Ссылка на видео
}
