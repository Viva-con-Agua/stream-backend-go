package database

import (
	"log"

	"../models"
	"../utils"
	"github.com/google/uuid"
)

/**
 * insert City into database
 */
func PostCity(r *models.CityCreate) (err error) {
	// Create uuid
	Uuid, err := uuid.NewRandom()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	// begin database query and handle error
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	// insert City
	_, err = tx.Exec("INSERT INTO City (uuid, name, pillar) VALUES(?, ?, ?)", Uuid.String(), r.Name, r.Pillar)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()
}

/**
 * select a list of models.Cities
 */
func GetCitiesList() (Cities []models.City, err error) {
	// Execute the Query
	rows, err := utils.DB.Query("SELECT * FROM City")
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	// define variable for each column
	var id int
	var uuid, name, pillar string
	// convert each row
	for rows.Next() {
		//scan row
		err = rows.Scan(&id, &uuid, &name, &pillar)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		// fill models.City
		City := new(models.City)
		City.Uuid = uuid
		City.Name = name
		City.Pillar = pillar
		Cities = append(Cities, *City)
	}
	return Cities, err
}

/**
 * select Cities for an given Crew_id
 */
func GetCitiesByCrewId(Crew_id int) (Cities []models.City, err error) {
	// Execute the Query
	rows, err := utils.DB.Query("SELECT * FROM Crew_has_City LEFT JOIN City ON Crew_has_City.City_Id = City.id WHERE Crew_has_City.Crew_id = ?", Crew_id)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	// define variable for each column
	var id int
	var uuid, name, pillar string
	// convert each row
	for rows.Next() {
		//scan row
		err = rows.Scan(&id, &id, &id, &uuid, &name, &pillar)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		// fill models.City
		City := new(models.City)
		City.Uuid = uuid
		City.Name = name
		City.Pillar = pillar
		Cities = append(Cities, *City)
	}
	return Cities, err
}
