package model

type Game struct {
	Code int16 `json:"code"`
	LPN  int16 `json:"lpn"`
}

type CreateGameRequest struct {
	Code int16 `json:"code"`
}
