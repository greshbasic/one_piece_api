package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	model "one-piece-api/models"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

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

	row := db.QueryRow("SELECT id, name, type, awakened, user_id FROM devil_fruits WHERE id = ?", id)
	if err := row.Scan(&df.ID, &df.Name, &df.Type, &df.Awakened, &df.UserID); err != nil {
		if err == sql.ErrNoRows {
			return df, fmt.Errorf("devilFruitByID %s: no such devil fruit", id)
		}
		return df, err
	}
	return df, nil
}

func allDevilFruits() ([]model.DevilFruit, error) {
	var dfs []model.DevilFruit

	rows, err := db.Query("SELECT id, name, type, awakened, user_id FROM devil_fruits")
	if err != nil {
		return nil, fmt.Errorf("devilFruits: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var df model.DevilFruit
		if err := rows.Scan(&df.ID, &df.Name, &df.Type, &df.Awakened, &df.UserID); err != nil {
			return nil, fmt.Errorf("devilFruits scan: %v", err)
		}
		dfs = append(dfs, df)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("devilFruits rows: %v", err)
	}

	return dfs, nil
}

func allCharacters() ([]model.Character, error) {
	var chars []model.Character

	rows, err := db.Query("SELECT id, name, age, status, race, origin, first_appearance, bounty, affiliation, haki, devil_fruit_id FROM characters")
	if err != nil {
		return nil, fmt.Errorf("characters: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Character
		var status sql.NullString
		var age sql.NullInt64
		var race, origin, firstAppearance, bounty, affiliation sql.NullString
		var hakiData sql.NullString
		var devilFruitID sql.NullString

		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&age,
			&status,
			&race,
			&origin,
			&firstAppearance,
			&bounty,
			&affiliation,
			&hakiData,
			&devilFruitID,
		); err != nil {
			return nil, fmt.Errorf("characters scan: %v", err)
		}

		// Handle nullable fields
		if age.Valid {
			ageInt := int(age.Int64)
			c.Age = &ageInt
		}
		if status.Valid {
			s := model.Status(status.String)
			c.Status = &s
		}
		if race.Valid {
			c.Race = &race.String
		}
		if origin.Valid {
			c.Origin = &origin.String
		}
		if firstAppearance.Valid {
			c.FirstAppearance = &firstAppearance.String
		}
		if bounty.Valid {
			c.Bounty = &bounty.String
		}
		if affiliation.Valid {
			c.Affiliation = &affiliation.String
		}

		if hakiData.Valid && hakiData.String != "" {
			var haki []model.HakiType
			if err := json.Unmarshal([]byte(hakiData.String), &haki); err == nil {
				c.Haki = haki
			}
		}

		if devilFruitID.Valid {
			c.DevilFruitID = &devilFruitID.String
		}

		chars = append(chars, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("characters rows: %v", err)
	}

	return chars, nil
}

func characterByID(id string) (model.Character, error) {
	var c model.Character
	var status string     // temporary for scanning status
	var hakiJSON []byte   // scan haki as bytes from DB
	var age sql.NullInt64 // nullable int

	row := db.QueryRow(`
		SELECT id, name, age, status, race, origin, first_appearance, bounty, affiliation, haki, devil_fruit_id 
		FROM characters WHERE id = ?`, id)

	err := row.Scan(
		&c.ID, &c.Name, &age, &status, &c.Race, &c.Origin, &c.FirstAppearance,
		&c.Bounty, &c.Affiliation, &hakiJSON, &c.DevilFruitID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, fmt.Errorf("characterByID %s: no such character", id)
		}
		return c, err
	}

	if age.Valid {
		ageInt := int(age.Int64)
		c.Age = &ageInt
	} else {
		c.Age = nil
	}

	s := model.Status(status)
	c.Status = &s

	if len(hakiJSON) > 0 {
		var haki []model.HakiType
		if err := json.Unmarshal(hakiJSON, &haki); err != nil {
			return c, fmt.Errorf("failed to unmarshal haki JSON: %v", err)
		}
		c.Haki = haki
	}

	return c, nil
}

func setupDB() *sql.DB {
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
