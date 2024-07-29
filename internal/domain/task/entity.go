package task

import (
	"time"
)

type Entity struct {
	ID       string    `db:"id" bson:"_id"`
	Title    string    `db:"title" bson:"title"`
	ActiveAt time.Time `db:"active_at" bson:"active_at"`
	Status   string    `db:"status" bson:"status"`
}

func (e *Entity) ParseToDayoffs() (*Entity, error) {
	if e.ActiveAt.Weekday() == time.Saturday || e.ActiveAt.Weekday() == time.Sunday {
		e.Title = "ВЫХОДНОЙ - " + e.Title
		return e, nil
	}
	return e, nil
}
