package model

type Material struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Cost int    `db:"cost"`
}
