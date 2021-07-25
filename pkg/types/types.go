package types

import "time"

type RadarProcessor interface {
	Process() *TechRadar
}

type TechRadar struct {
	Entries []Entry  `json:"entries"`
	Quadrants [] Quadrant `json:"quadrants"`
	Rings []Ring  `json:"rings"`
}
type Quadrant struct {
		ID   string `json:"id"`
		Name string `json:"name"`
}
type Ring struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
type Entry struct {
	Timeline []TimelineEntry `json:"timeline"`
	URL         string `json:"url"`
	Key         string `json:"key"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Quadrant    string `json:"quadrant"`
	Description string `json:"description,omitempty"`
}
type TimelineEntry struct {
Moved       int       `json:"moved"`
RingID      string    `json:"ringId"`
Date        time.Time `json:"date"`
Description string    `json:"description"`
}
