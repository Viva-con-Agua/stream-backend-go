package database

import (
	"log"
	"time"

	"../models"
	"../utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

/**
 * select Supporter
 */
func GetSupporter(search string) (Supporters []models.Supporter, err error) {
	// execute the query
	SupporterQuery := "SELECT Supporter.id, Supporter.uuid, Profile.email, Profile.first_name, Profile.last_name, Supporter.updated, Supporter.created " +
		"FROM Supporter LEFT JOIN Profile ON Supporter.id = Profile.Supporter_id " +
		"LEFT JOIN Supporter_has_Role ON Supporter.id = Supporter_has_Role.Supporter_id " +
		"LEFT JOIN Role ON Supporter_has_Role.Role_Id = Role.id " +
		"WHERE Supporter.uuid = ? " +
		"GROUP BY Supporter.id "
	rows, err := utils.DB.Query(SupporterQuery, search)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	//define variable for Supporter database id
	var id int
	Supporter := new(models.Supporter)
	for rows.Next() {
		//scan row
		err = rows.Scan(&id, &Supporter.Uuid, &Supporter.Email, &Supporter.FirstName, &Supporter.LastName, &Supporter.Updated, &Supporter.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		log.Print(id)
		// get roles by id
		var roles []models.Role
		roles, err = GetRolesBySupporterId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		//join roles to Supporter
		Supporter.Roles = roles
	}
	if id == 0 {
		return nil, err
	}
	Supporters = append(Supporters, *Supporter)
	return Supporters, err
}

/**
 * select list of Supporter
 */
func GetSupporterList(page *models.Page, sort string, filter *models.FilterSupporter) (Supporters []models.Supporter, err error) {
	// execute the query
	SupporterQuery := "SELECT u.id, u.uuid, p.email, p.first_name, p.last_name, u.updated, u.created " +
		"FROM Supporter AS u LEFT JOIN Profile AS p ON u.id = p.Supporter_id " +
		"LEFT JOIN Supporter_has_Role ON u.id = Supporter_has_Role.Supporter_id " +
		"LEFT JOIN Role ON Supporter_has_Role.Role_Id = Role.id " +
		"WHERE p.email LIKE ? " +
		"GROUP BY u.id " +
		sort + " " +
		"LIMIT ?, ?"
	rows, err := utils.DB.Query(SupporterQuery, filter.Email, page.Offset, page.Count)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	// variable for Supporter database id
	var id int
	// convert each row
	for rows.Next() {
		//create Supporter
		Supporter := new(models.Supporter)
		//scan row and fill Supporter
		err = rows.Scan(&id, &Supporter.Uuid, &Supporter.Email, &Supporter.FirstName, &Supporter.LastName, &Supporter.Updated, &Supporter.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		// get roles by id
		var roles []models.Role
		roles, err = GetRolesBySupporterId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}

		// join roles to Supporter
		Supporter.Roles = roles
		// append to list of Supporter
		Supporters = append(Supporters, *Supporter)
	}
	return Supporters, err
}

/**
 * update Supporter
 */
func UpdateSupporter(Supporter *models.Supporter) (err error) {
	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//slect id
	rows, err := tx.Query("SELECT id FROM Supporter WHERE uuid = ?", Supporter.Uuid)
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Print("Database Error: ", err)
			return err
		}
	}
	//if id == 0 return NotFound
	if id == 0 {
		err = utils.ErrorNotFound
		return err
	}

	//update Supporter Supporter
	_, err = tx.Exec("UPDATE Supporter SET updated = ? WHERE id = ?", time.Now().Unix(), id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	//update profile
	_, err = tx.Exec("UPDATE Profile SET first_name = ?, last_name = ? WHERE Supporter_id = ?", Supporter.FirstName, Supporter.LastName, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

/**
 * Create Supporter
 */
func CreateSupporter(Supporter *models.SupporterCreate) (err error) {
	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}

	//Insert Supporter Supporter
	id := uuid.New()
	_, err = tx.Exec("UPDATE Supporter SET updated = ? WHERE id = ?", time.Now().Unix(), id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	//update profile
	_, err = tx.Exec("UPDATE Profile SET first_name = ?, last_name = ? WHERE Supporter_id = ?", Supporter.FirstName, Supporter.LastName, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

func DeleteSupporter(deleteBody *models.DeleteBody) (err error) {
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//slect id
	rows, err := tx.Query("SELECT id FROM Supporter WHERE uuid = ?", deleteBody.Uuid)
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Print("Database Error: ", err)
			return err
		}
	}
	//if id == 0 return NotFound
	if id == 0 {
		err = utils.ErrorNotFound
		return err
	}

	//update Supporter Supporter
	_, err = tx.Exec("DELETE FROM Supporter WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

/**
 * join Supporter and role entry via Supporter_has_Role table
 */
func JoinSupporterRole(assign *models.AssignBody) (err error) {
	// select Supporter_id from database
	rows, err := utils.DB.Query("SELECT id FROM Supporter WHERE uuid = ?", assign.Assign)
	if err != nil {
		log.Print("Database Error", err)
		return err
	}
	// select Supporter_id from rows
	var SupporterId int
	for rows.Next() {
		err = rows.Scan(&SupporterId)
		if err != nil {
			log.Print("Database Error: ", err)
			return err
		}
	}
	// select role_id from database
	rows2, err := utils.DB.Query("SELECT id FROM Role WHERE uuid = ?", assign.To)
	if err != nil {
		log.Print("Database Error", err)
		return err
	}
	//select Supporter_id from rows
	var roleId int
	for rows2.Next() {
		err = rows2.Scan(&roleId)
		if err != nil {
			log.Print("Database Error: ", err)
			return err
		}
	}
	// begin database query and handle error
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	// insert Supporter_has_Role
	_, err = tx.Exec("INSERT INTO Supporter_has_Role (Supporter_id, Role_Id) VALUES(?, ?)", SupporterId, roleId)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()
}
