package models

import "time"

type Course struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Urls              []string  `json:"urls"`
	Priority          int       `json:"priority"` // 1- > low, 2 -> medium, 3 -> high
	Checkoff          int       `json:"checkoff"` //1 -> daily, 2 -> weekly, 3 -> monthly
	CheckoffTimeStamp time.Time `json:"checkoff_time"`
	Notes             []string
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
