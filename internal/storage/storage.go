package storage

import (
	"encoding/json"

	expression "github.com/Aspikk/Distributed-Calculator/internal/entities/expression"
	queue "github.com/Aspikk/Distributed-Calculator/internal/entities/queue"
	task "github.com/Aspikk/Distributed-Calculator/internal/entities/task"
)

type DataBase struct {
	Expressions *[]*expression.Expression `json:"expressions"`
	Tasks       *queue.Queue[*task.Task]
}

func (db *DataBase) MarshallJSON() ([]byte, error) {
	var tmp struct {
		Expressions *[]*expression.Expression `json:"expressions"`
	}

	tmp.Expressions = db.Expressions
	return json.Marshal(&tmp)
}

var DB *DataBase

func init() {
	DB = &DataBase{
		Expressions: &[]*expression.Expression{},
		Tasks:       queue.New[*task.Task](),
	}
}
