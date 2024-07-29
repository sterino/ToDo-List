package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"todo-list/internal/domain/task"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) Create(ctx context.Context, data task.Entity) (id string, err error) {
	query := `INSERT INTO tasks (title, active_at) VALUES ($1,$2) RETURNING id;`
	args := []any{
		data.Title,
		data.ActiveAt,
	}
	if err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
		}
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return "", task.ErrorNotFound
		}
	}
	return
}

func (tr *TaskRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM tasks WHERE id = $1 RETURNING id;`
	args := []any{id}
	if err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
			return
		}
	}

	return
}

func (tr *TaskRepository) Get(ctx context.Context, id string) (dest task.Entity, err error) {
	query := `SELECT * FROM tasks WHERE id = $1;`
	args := []any{id}
	err = tr.db.GetContext(ctx, &dest, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		err = task.ErrorNotFound
	}
	return
}

func (tr *TaskRepository) List(ctx context.Context, status string) (dest []task.Entity, err error) {
	query := `SELECT * FROM tasks WHERE status = $1 AND active_at <= CURRENT_DATE ORDER BY active_at;`
	err = tr.db.SelectContext(ctx, &dest, query, status)
	if err != nil {
		return
	}
	if len(dest) == 0 {
		err = task.ErrorNotFound
		return
	}
	return
}

func (tr *TaskRepository) Update(ctx context.Context, id string, data task.Entity) (err error) {
	query := `UPDATE tasks SET title = $1, active_at = $2 WHERE id = $3 RETURNING id;`
	args := []any{
		data.Title,
		data.ActiveAt,
		id,
	}
	if err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
		}
	}
	return
}

func (tr *TaskRepository) Done(ctx context.Context, id string) (err error) {
	query := `UPDATE tasks SET status = 'done' WHERE id = $1 RETURNING id;`
	args := []any{id}
	if err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
			return
		}
	}
	return
}
