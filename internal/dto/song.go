package dto

type SongCreateRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
