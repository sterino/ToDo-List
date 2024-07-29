package task

import (
	"errors"
	"time"
)

var (
	ErrorNotFound      = errors.New("task not found")
	ErrorInvalidTitle  = errors.New("invalid title")
	ErrorInvalidDate   = errors.New("invalid date format")
	ErrorInvalidStatus = errors.New("invalid status")
)

type Request struct {
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
}

type Response struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"active_at"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:       entity.ID,
		Title:    entity.Title,
		ActiveAt: entity.ActiveAt,
	}
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, entity := range data {
		res = append(res, ParseFromEntity(entity))
	}
	return
}

func (r *Request) Validate() error {
	if r.Title == "" && len(r.Title) < 200 {
		return ErrorInvalidTitle
	}
	if _, err := time.Parse("2006-01-02", r.ActiveAt); err != nil {
		return ErrorInvalidDate
	}
	return nil
}

func IsValidStatus(priority string) bool {
	validPriorities := map[string]bool{
		"active": true,
		"done":   true,
		"":       true,
	}
	return validPriorities[priority]
}

func ParseDate(date string) (data time.Time) {
	data, _ = time.Parse("2006-01-02", date)
	return
}
