package model

type Log struct {
	Medicines []string `json:"medicines" bson:"medicines"`
	Meta      `bson:",inline"`
}
