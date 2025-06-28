package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"one-piece-api/models"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getDevilFruitByID(id int64) (models.DevilFruit, error) {
	var df models.DevilFruit
	var userID sql.NullInt64

	query := `
        SELECT id, name, type, awakened, user_id
        FROM devil_fruits
        WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&df.ID, &df.Name, &df.Type, &df.Awakened, &userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return df, fmt.Errorf("devil fruit %d not found", id)
		}
		return df, err
	}

	if userID.Valid {
		df.UserID = &userID.Int64
	} else {
		df.UserID = nil
	}

	return df, nil
}

func getCharacterByID(id int64) (models.Character, error) {
	var c models.Character

	var bounty sql.NullInt64
	var hakiBytes sql.NullString
	var affiliation, origin, status sql.NullString
	var age sql.NullInt64

	var dfID sql.NullInt64
	var dfName, dfType sql.NullString
	var dfAwakened sql.NullBool

	query := `
        SELECT
            c.id, c.name, c.bounty, c.haki, c.affiliation, c.origin, c.status, c.age,
            d.id, d.name, d.type, d.awakened
        FROM characters c
        LEFT JOIN devil_fruits d ON d.user_id = c.id
        WHERE c.id = ?`

	row := db.QueryRow(query, id)
	err := row.Scan(
		&c.ID, &c.Name, &bounty, &hakiBytes, &affiliation, &origin, &status, &age,
		&dfID, &dfName, &dfType, &dfAwakened,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, fmt.Errorf("character %d not found", id)
		}
		return c, err
	}

	if bounty.Valid {
		c.Bounty = &bounty.Int64
	}

	if hakiBytes.Valid {
		var hakiMap models.Haki
		if err := json.Unmarshal([]byte(hakiBytes.String), &hakiMap); err == nil {
			c.Haki = &hakiMap
		}
	}

	if affiliation.Valid {
		c.Affiliation = &affiliation.String
	}
	if origin.Valid {
		c.Origin = &origin.String
	}
	if status.Valid {
		c.Status = &status.String
	}
	if age.Valid {
		ageInt := int(age.Int64)
		c.Age = &ageInt
	}

	if dfID.Valid {
		c.DevilFruit = &models.DevilFruit{
			ID:       dfID.Int64,
			Name:     dfName.String,
			Type:     dfType.String,
			Awakened: dfAwakened.Bool,
		}
	}

	return c, nil
}

func getAllCharacters() ([]models.Character, error) {
	var characters []models.Character

	query := `
        SELECT
            c.id, c.name, c.bounty, c.haki, c.affiliation, c.origin, c.status, c.age,
            d.id, d.name, d.type, d.awakened
        FROM characters c
        LEFT JOIN devil_fruits d ON d.user_id = c.id
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Character

		var bounty sql.NullInt64
		var hakiBytes sql.NullString
		var affiliation, origin, status sql.NullString
		var age sql.NullInt64

		var dfID sql.NullInt64
		var dfName, dfType sql.NullString
		var dfAwakened sql.NullBool

		err := rows.Scan(
			&c.ID, &c.Name, &bounty, &hakiBytes, &affiliation, &origin, &status, &age,
			&dfID, &dfName, &dfType, &dfAwakened,
		)
		if err != nil {
			return nil, err
		}

		if bounty.Valid {
			c.Bounty = &bounty.Int64
		}
		if hakiBytes.Valid {
			var hakiMap models.Haki
			if err := json.Unmarshal([]byte(hakiBytes.String), &hakiMap); err == nil {
				c.Haki = &hakiMap
			}
		}
		if affiliation.Valid {
			c.Affiliation = &affiliation.String
		}
		if origin.Valid {
			c.Origin = &origin.String
		}
		if status.Valid {
			c.Status = &status.String
		}
		if age.Valid {
			ageInt := int(age.Int64)
			c.Age = &ageInt
		}

		if dfID.Valid {
			c.DevilFruit = &models.DevilFruit{
				ID:       dfID.Int64,
				Name:     dfName.String,
				Type:     dfType.String,
				Awakened: dfAwakened.Bool,
			}
		}

		characters = append(characters, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return characters, nil
}

func getAllDevilFruits() ([]models.DevilFruit, error) {
	var devilFruits []models.DevilFruit

	query := `
        SELECT id, name, type, awakened, user_id
        FROM devil_fruits
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var df models.DevilFruit
		var userID sql.NullInt64

		err := rows.Scan(&df.ID, &df.Name, &df.Type, &df.Awakened, &userID)
		if err != nil {
			return nil, err
		}

		if userID.Valid {
			df.UserID = &userID.Int64
		} else {
			df.UserID = nil
		}

		devilFruits = append(devilFruits, df)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devilFruits, nil
}

func setupDB() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "bugs"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "one_piece_data"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Punk Records DB!")
	return db
}
