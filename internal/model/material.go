package model

type Material struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Cost int    `db:"cost" json:"cost"`
}
