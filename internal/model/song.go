package model

const TBL_NAME = "songs"

type Song struct {
	ID          int    `db:"id"`
	Group       string `db:"group_song"`
	Song        string `db:"song"`
	ReleaseDate string `db:"release_date"`
	Text        string `db:"text"`
	Link        string `db:"link"`
}
