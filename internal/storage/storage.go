package storage

import entities "github.com/Aspikk/Distributed-Calculator/internal/entities/expression"

type DataBase struct {
	Expressions []*entities.Expression `json:"expressions"`
}

var DB *DataBase

func init() {
	DB = &DataBase{
		Expressions: make([]*entities.Expression, 0),
	}
}
