package models

import (
	"fmt"
	"time"

	"haseeb.khan/event-booking/database"
)

type Event struct {
	ID          int64 //`json:"-"`
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int64
}

func (event *Event) Save() error {
	/* query := `INSERT INTO events(name, description, location, date_time, user_id) VALUES($1, $2, $3, $4, $5)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id
	return err */

	query := `
	INSERT INTO events(name, description, location, date_time, user_id) 
	VALUES($1, $2, $3, $4, $5) 
	RETURNING id
	`

	var id int64
	err := database.DB.QueryRow(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println("ID Returned: ", id)
	event.ID = id
	return nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := database.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = $1, description = $2, location = $3, date_time = $4
	WHERE id = $5
	`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = $1"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (event Event) Register(uId int64) error {
	query := "INSERT INTO registrations(user_id, event_id) VALUES($1,$2)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uId, event.ID)
	return err
}

func (event Event) Unregister(uId int64) error {
	query := "DELETE FROM registrations WHERE event_id = $1 AND user_id = $2"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, uId)
	return err
}
