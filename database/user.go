package database

import (
	"log"

	"../models"
	"../utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func InsertUser(u *models.UserCreate) (err error) {
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

	//insert user
	res, err := tx.Exec("INSERT INTO User (uuid, first_name, last_name) VALUES(?, ?, ?)", Uuid.String(), u.FirstName, u.LastName)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}

	res, err = tx.Exec("INSERT INTO LoginInfo (email, password, User_id) VALUES(?, ?, ?)", u.LoginInfo.Email, u.LoginInfo.Password, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}
