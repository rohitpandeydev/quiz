package models

type Question struct {
	Question string   `json:"question"`
	Answer   []string `json:"answer"`
}

type Score struct {
	Correct int
	Total   int
}
