package server

type Server struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Alive bool
}
