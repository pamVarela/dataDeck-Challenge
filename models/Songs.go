package models

type Songs struct {
	ID     int    `json:"id"`
	Artist string `json:"artist"`
	Song   string `json:"song"`   
	Genre  string `json:"genre"`
	Length int    `json:"length"`
}
