package models

import (
	"time"

	"phpguru.net/go-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
        INSERT INTO events(name, description, location, dateTime, user_id)
        VALUES(?, ?, ?, ?, ?)
    `
	stmt, err := db.GetDB().Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := rs.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents(userId int64) ([]Event, error) {
	var events []Event

	query := "SELECT * from events where user_id = ?"
	rows, err := db.GetDB().Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {

	query := "SELECT * from events where id = ?"
	row := db.GetDB().QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, err
}

func (event *Event) Update() error {

	query := `
        UPDATE events
        SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
        WHERE id = ?
    `
	stmt, err := db.GetDB().Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEventById(id int64) error {
	query := "DELETE from events where id = ?"
	stmt, err := db.GetDB().Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	// rs.RowsAffected()
	return err
}
