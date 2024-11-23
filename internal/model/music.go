package model

type Music struct {
	ID          int    `db:"id"`
	Group       string `db:"group"`
	Song        string `db:"song"`
	ReleaseDate string `db:"releaseDate"`
	Text        string `db:"text"`
	Link        string `db:"link"`
}
