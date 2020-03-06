package database

import (
	"../models"
	"../utils"
	"github.com/google/uuid"
	"log"
	"time"
)

/**
 * select Supporter
 */
func GetSupporter(search string) (Supporters []models.Supporter, err error) {
	// execute the query
	SupporterQuery := "SELECT profile.id, profile.uuid, profile.email, profile.firstname, profile.lastname, CONCAT(profile.firstname, ' ', profile.lastname) AS fullname, profile.birthdate, profile.sex, profile.updated, profile.created " +
		"FROM profile LEFT JOIN profile_has_address ON profile.id = profile_has_address.profile_id " +
		"LEFT JOIN avatar ON profile.id = avatar.profile_id " +
		"LEFT JOIN profile_has_entity ON profile.id = profile_has_entity.profile_id " +
		"WHERE profile.uuid = ? " +
		"GROUP BY profile.id "
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
	SupporterQuery := "SELECT profile.id, profile.uuid, profile.email, profile.firstname, profile.lastname, CONCAT(profile.firstname, ' ', profile.lastname) AS fullname, profile.updated, profile.created " +
		"FROM profile LEFT JOIN profile_has_entity ON profile.id = profile_has_entity.profile_id " +
		"WHERE profile.email LIKE ? " +
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
	rows, err := tx.Query("SELECT id FROM profile WHERE uuid = ?", Supporter.Uuid)
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
	_, err = tx.Exec("UPDATE profile SET updated = ? WHERE id = ?", time.Now().Unix(), id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	//update profile
	_, err = tx.Exec("UPDATE profile SET firstname = ?, lastname = ? WHERE id = ?", Supporter.FirstName, Supporter.LastName, id)
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

	// Check for existing profile
	rows, err := tx.Query("SELECT id FROM profile WHERE email = ?", Supporter.email)
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Print("Database Error: ", err)
			return err
		}
	}
	//if id == 0 return NotFound
	if id != 0 {
		err = utils.ErrorConflict
		return err
	}

	// Insert Profile
	id := uuid.New()
	_, err = tx.Exec("INSERT INTO profile (uuid, firstname, lastname, email, mobile, birthdate, sex, updated, created) VALUES "+
		"(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		Supporter.uuid, Supporter.firstname, Supporter.lastname, Supporter.email, Supporter.mobile, Supporter.birthdate, time.Now().Unix(), time.Now().Unix())
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

	// Delete profike
	// TODO DELETE PROFILE AND CORRESPONDING RELATIONS
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
