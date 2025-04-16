package model

import "time"

type Questionnaire struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
