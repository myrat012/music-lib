package dto

import (
	"strconv"
)

type SongCreateRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongRequest struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

type SongGetRequest struct {
	Group string
	Song  string
	Page  int
	Limit int
}

func (sgr *SongGetRequest) ToStruct(group, song, page, limit string) error {
	p, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}

	sgr.Group = group
	sgr.Song = song
	sgr.Page = p
	sgr.Limit = l

	return nil
}
