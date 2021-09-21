package dto

type League struct {
	id         uint64 `json:"id" form:"id"`
	Date       string `json:"date" form:"date" binding:"required"`
	HomeTeam   string `json:"home_team" form:"home_team" binding:"required"`
	AwayTeam   string `json:"away_team" form:"away_team" binding:"required"`
	FTHG       uint64 `json:"fthg" form:"fthg" binding:"required"  validate:"max:5"`
	FTAG       uint64 `json:"ftag" form:"ftag" binding:"required"  validate:"min:1"`
	Referee    string `json:"referee" form:"referee"`
	leagueName string `json:"leage_name" form:"league_name" binding:"required"`
}
