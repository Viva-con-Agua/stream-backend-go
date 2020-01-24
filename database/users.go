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
 * insert user into database
 */
func SignUp(u *models.UserCreate) (err error) {
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

	//insert user
	res, err := tx.Exec("INSERT INTO User (uuid, updated, created) VALUES(?, ?, ?)", Uuid.String(), time.Now().Unix(), time.Now().Unix())
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	// get user id
	id, err := res.LastInsertId()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	// insert loginInfo
	res, err = tx.Exec("INSERT INTO LoginInfo (email, password, User_id) VALUES(?, ?, ?)", u.LoginInfo.Email, u.LoginInfo.Password, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	// insert profile
	res, err = tx.Exec("INSERT INTO Profile (email, first_name, last_name, User_id) VALUES(?, ?, ?, ?)", u.LoginInfo.Email, u.FirstName, u.LastName, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

/**
 * select user
 */
func GetUser(search string) (users []models.User, err error) {
	// execute the query
	userQuery := "SELECT User.id, User.uuid, Profile.email, Profile.first_name, Profile.last_name, User.updated, User.created " +
		"FROM User LEFT JOIN Profile ON User.id = Profile.User_id " +
		"LEFT JOIN User_has_Role ON User.id = User_has_Role.User_id " +
		"LEFT JOIN Role ON User_has_Role.Role_Id = Role.id " +
		"WHERE User.uuid = ? " +
		"GROUP BY User.id "
	rows, err := utils.DB.Query(userQuery, search)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	//define variable for user database id
	var id int
	user := new(models.User)
	for rows.Next() {
		//scan row
		err = rows.Scan(&id, &user.Uuid, &user.Email, &user.FirstName, &user.LastName, &user.Updated, &user.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		log.Print(id)
		// get roles by id
		var roles []models.Role
		roles, err = GetRolesByUserId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		//join roles to user
		user.Roles = roles
	}
	if id == 0 {
		return nil, err
	}
	users = append(users, *user)
	return users, err
}

/**
 * select list of user
 */
func GetUserList(page *models.Page, sort string, filter *models.FilterUser) (users []models.User, err error) {
	// execute the query
	userQuery := "SELECT u.id, u.uuid, p.email, p.first_name, p.last_name, u.updated, u.created " +
		"FROM User AS u LEFT JOIN Profile AS p ON u.id = p.User_id " +
		"LEFT JOIN User_has_Role ON u.id = User_has_Role.User_id " +
		"LEFT JOIN Role ON User_has_Role.Role_Id = Role.id " +
		"WHERE p.email LIKE ? " +
		"GROUP BY u.id " +
		sort + " " +
		"LIMIT ?, ?"
	rows, err := utils.DB.Query(userQuery, filter.Email, page.Offset, page.Count)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	// variable for user database id
	var id int
	// convert each row
	for rows.Next() {
		//create user
		user := new(models.User)
		//scan row and fill user
		err = rows.Scan(&id, &user.Uuid, &user.Email, &user.FirstName, &user.LastName, &user.Updated, &user.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		// get roles by id
		var roles []models.Role
		roles, err = GetRolesByUserId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}

		// join roles to user
		user.Roles = roles
		// append to list of user
		users = append(users, *user)
	}
	return users, err
}

/**
 * update user
 */
func UpdateUser(user *models.User) (err error) {
	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//slect id
	rows, err := tx.Query("SELECT id FROM User WHERE uuid = ?", user.Uuid)
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

	//update user user
	_, err = tx.Exec("UPDATE User SET updated = ? WHERE id = ?", time.Now().Unix(), id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	//update profile
	_, err = tx.Exec("UPDATE Profile SET first_name = ?, last_name = ? WHERE User_id = ?", user.FirstName, user.LastName, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

func DeleteUser(deleteBody *models.DeleteBody) (err error) {
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//slect id
	rows, err := tx.Query("SELECT id FROM User WHERE uuid = ?", deleteBody.Uuid)
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

	//update user user
	_, err = tx.Exec("DELETE FROM User WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

/**
 * join user and role entry via User_has_Role table
 */
func JoinUserRole(assign *models.AssignBody) (err error) {
	// select user_id from database
	rows, err := utils.DB.Query("SELECT id FROM User WHERE uuid = ?", assign.Assign)
	if err != nil {
		log.Print("Database Error", err)
		return err
	}
	// select user_id from rows
	var userId int
	for rows.Next() {
		err = rows.Scan(&userId)
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
	//select user_id from rows
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
	// insert User_has_Role
	_, err = tx.Exec("INSERT INTO User_has_Role (User_id, Role_Id) VALUES(?, ?)", userId, roleId)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()
}
