package main

import (
	"database/sql"
	"fmt"
	"log"

	model "one-piece-api/models"

	"github.com/go-sql-driver/mysql"
)

// var dbCharacters *sql.DB
var dbDevilFruits *sql.DB

// func characterByID(id int64) (model.Character, error) {
// 	var c model.Character

// 	row := dbCharacters.QueryRow("SELECT * FROM character WHERE id = ?", id)
// 	if err := row.Scan(&c.ID, &c.Name, &c.Age, &c.Status, &c.Race, &c.Origin, &c.FirstAppearance, &c.Bounty, &c.Affiliation, &c.Haki, &c.DevilFruit); err != nil {
// 		if err == sql.ErrNoRows {
// 			return c, fmt.Errorf("albumsById %d: no such album", id)
// 		}
// 		return c, fmt.Errorf("albumsById %d: %v", id, err)
// 	}
// 	return c, nil
// }

func devilFruitByID(id string) (model.DevilFruit, error) {
	var df model.DevilFruit

	row := dbDevilFruits.QueryRow("SELECT id, name, type, awakened, user_id FROM devil_fruits WHERE id = ?", id)
	if err := row.Scan(&df.ID, &df.Name, &df.Type, &df.Awakened, &df.UserID); err != nil {
		if err == sql.ErrNoRows {
			return df, fmt.Errorf("devilFruitByID %s: no such devil fruit", id)
		}
		return df, fmt.Errorf("devilFruitByID %s: %v", id, err)
	}
	return df, nil
}

func setupCharactersDB() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "bugs"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "characters"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Characters DB!")
	return db
}

func setupDevilFruitsDB() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "bugs"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "devil_fruits"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Devil Fruits DB!")
	return db
}
