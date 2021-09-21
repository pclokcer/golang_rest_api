package entity

type League struct {
	id         uint64 `gorm: "primary_key:auto_increment" json:"id"`
	Date       string `gorm: "type: Datetime" json:"date"`
	HomeTeam   string `gorm: "type: varchar(100)" json:"home_team"`
	AwayTeam   string `gorm: "type: varchar(100)" json:"away_team"`
	FTHG       uint64 `gorm: "type: int" json:"fthg"`
	FTAG       uint64 `gorm: "type: int" json:"ftag"`
	Referee    string `gorm: "type: varchar(45)" json:"referee"`
	leagueName string `gorm: "type: varchar(45)" json:"leage_name"`
}
