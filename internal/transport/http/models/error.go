package models

type Error struct {
	Code    int    `json:"code" binding:"required" example:"1"`
	Message string `json:"message" binding:"required" example:"internal server error"`
}
