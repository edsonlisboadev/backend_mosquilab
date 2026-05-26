package repository

import (
	"database/sql"
	"mosquilab/internal/models"
	"time"
)

type EventRepo struct{ db *sql.DB }

func NewEventRepo(db *sql.DB) *EventRepo { return &EventRepo{db: db} }

func (r *EventRepo) ListFuture() ([]models.Event, error) {
	rows, err := r.db.Query(`
		SELECT id, title, COALESCE(description,''), COALESCE(location,''),
		       event_date::text, event_time::text, COALESCE(event_type,''), created_at, updated_at
		FROM events
		WHERE event_date >= CURRENT_DATE
		ORDER BY event_date, event_time`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanEvents(rows)
}

func (r *EventRepo) ListAll() ([]models.Event, error) {
	rows, err := r.db.Query(`
		SELECT id, title, COALESCE(description,''), COALESCE(location,''),
		       event_date::text, event_time::text, COALESCE(event_type,''), created_at, updated_at
		FROM events ORDER BY event_date DESC, event_time DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanEvents(rows)
}

func scanEvents(rows *sql.Rows) ([]models.Event, error) {
	var events []models.Event
	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location,
			&e.EventDate, &e.EventTime, &e.EventType, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	if events == nil {
		events = []models.Event{}
	}
	return events, nil
}

func (r *EventRepo) Create(e *models.Event) (*models.Event, error) {
	var created models.Event
	err := r.db.QueryRow(`
		INSERT INTO events (title, description, location, event_date, event_time, event_type)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, title, COALESCE(description,''), COALESCE(location,''),
		          event_date::text, event_time::text, COALESCE(event_type,''), created_at, updated_at`,
		e.Title, e.Description, e.Location, e.EventDate, e.EventTime, e.EventType,
	).Scan(&created.ID, &created.Title, &created.Description, &created.Location,
		&created.EventDate, &created.EventTime, &created.EventType, &created.CreatedAt, &created.UpdatedAt)
	return &created, err
}

func (r *EventRepo) Update(id int, e *models.Event) (*models.Event, error) {
	var updated models.Event
	err := r.db.QueryRow(`
		UPDATE events SET title=$1, description=$2, location=$3,
		                  event_date=$4, event_time=$5, event_type=$6, updated_at=$7
		WHERE id=$8
		RETURNING id, title, COALESCE(description,''), COALESCE(location,''),
		          event_date::text, event_time::text, COALESCE(event_type,''), created_at, updated_at`,
		e.Title, e.Description, e.Location, e.EventDate, e.EventTime, e.EventType, time.Now(), id,
	).Scan(&updated.ID, &updated.Title, &updated.Description, &updated.Location,
		&updated.EventDate, &updated.EventTime, &updated.EventType, &updated.CreatedAt, &updated.UpdatedAt)
	return &updated, err
}

func (r *EventRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM events WHERE id=$1`, id)
	return err
}
