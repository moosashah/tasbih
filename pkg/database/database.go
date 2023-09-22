package database

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type TasbihDB struct {
	db      *sql.DB
	dataDir string
}

type person struct {
	ID      string
	count   int
	room_id string
}

type room struct {
	ID    string
	count int
}

func (t *TasbihDB) createRooms() error {
	_, err := t.db.Exec(`CREATE TABLE IF NOT EXISTS rooms (
		id TEXT PRIMARY KEY,
		count INTEGER,
	)`)
	return err
}

func (t *TasbihDB) createPeople() error {
	_, err := t.db.Exec(`CREATE TABLE IF NOT EXISTS people (
		id TEXT PRIMARY KEY,
		count INTEGER,
		room_id TEXT,
		FOREIGN KEY (room_id) REFERENCES rooms(id)
	)`)
	return err
}

func (t *TasbihDB) addRoom(person_id string) error {
	room := room{
		ID:    uuid.New().String(),
		count: 0,
	}
	_, err := t.db.Exec(
		"INSERT INTO rooms (id,count,person_id) VALUES(?,?,?)",
		room.ID, room.count,
		person_id,
	)
	return err
}

func (t *TasbihDB) addPerson() error {
	count := 0
	_, err := t.db.Exec(
		"INSERT INTO people (count) VALUES(?)",
		count,
	)
	return err
}

func (t *TasbihDB) addPersonToRoom(room_id, person_id string) error {
	// update person to add room id
	_, err := t.db.Exec(`insert into people ()`)
	return err
}

func InitDB() (*TasbihDB, error) {
	path := "tmp/db/database.db"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	t := TasbihDB{
		db, path,
	}
	if err := t.createPeople(); err != nil {
		return nil, err
	}
	if err := t.createRooms(); err != nil {
		return nil, err
	}
	return &t, nil
}
