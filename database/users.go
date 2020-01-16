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
 * select list of user
 */
func GetUserList() (users []models.User, err error) {
	// execute the query
	rows, err := utils.DB.Query("SELECT * FROM User LEFT JOIN Profile ON User.id = Profile.User_id ")
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	// define variable for each column
	var id, profile_id, user_id, updated, created int
	var uuid, email, first_name, last_name string
	// convert each row
	for rows.Next() {
		//scan row
		err = rows.Scan(&id, &uuid, &updated, &created, &profile_id, &email, &first_name, &last_name, &user_id)
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

		// fill models.User
		user := new(models.User)
		user.Uuid = uuid
		user.Email = email
		user.FirstName = first_name
		user.LastName = last_name
		user.Updated = updated
		user.Created = created
		user.Roles = roles
		users = append(users, *user)
	}
	return users, err
}

/**
 * join user and role entry via User_has_Role table
 */
func JoinUserRole(userRole *models.UserRole) (err error) {
	// select user_id from database
	rows, err := utils.DB.Query("SELECT id FROM User WHERE uuid = ?", userRole.UserId)
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
	rows2, err := utils.DB.Query("SELECT id FROM Role WHERE uuid = ?", userRole.RoleId)
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
