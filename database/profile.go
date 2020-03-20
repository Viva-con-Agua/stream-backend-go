package database

import (
	"../models"
	"../utils"
	"github.com/google/uuid"
	"log"
	"time"
)

/**
 * select Profile
 */
func GetProfile(search string) (Profiles []models.Profile, err error) {
	// execute the query
	ProfileQuery := "SELECT profile.id, profile.uuid, profile.email, profile.firstname, profile.lastname, CONCAT(profile.firstname, ' ', profile.lastname) AS fullname, profile.birthdate, profile.sex, profile.updated, profile.created " +
		"FROM profile LEFT JOIN profile_has_address ON profile.id = profile_has_address.profile_id " +
		"LEFT JOIN avatar ON profile.id = avatar.profile_id " +
		"LEFT JOIN profile_has_entity ON profile.id = profile_has_entity.profile_id " +
		"WHERE profile.uuid = ? " +
		"GROUP BY profile.id "
	rows, err := utils.DB.Query(ProfileQuery, search)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	//define variable for Profile database id
	var id int
	Profile := new(models.Profile)
	for rows.Next() {
		//scan row
		err = rows.Scan(&id, &Profile.Uuid, &Profile.Email, &Profile.FirstName, &Profile.LastName, &Profile.FullName, &Profile.Birthdate, &Profile.Sex, &Profile.Updated, &Profile.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		log.Print(id)
		// get roles by id
	}
	if id == 0 {
		return nil, err
	}
	Profiles = append(Profiles, *Profile)
	return Profiles, err
}

/**
 * select list of Profile
 */
func GetProfileList(page *models.Page, sort string, filter *models.FilterProfile) (Profiles []models.Profile, err error) {
	// execute the query
	ProfileQuery := "SELECT profile.id, profile.uuid, profile.email, profile.firstname, profile.lastname, CONCAT(profile.firstname, ' ', profile.lastname) AS fullname, profile.mobile, profile.birthdate, profile.sex, profile.updated, profile.created " +
		"FROM profile LEFT JOIN profile_has_entity ON profile.id = profile_has_entity.profile_id " +
		"WHERE profile.email LIKE ? " +
		"GROUP BY profile.id " +
		sort + " " +
		"LIMIT ?, ?"
	rows, err := utils.DB.Query(ProfileQuery, filter.Email, page.Offset, page.Count)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}
	// variable for Profile database id
	var id int
	// convert each row
	for rows.Next() {
		//create Profile
		Profile := new(models.Profile)
		//scan row and fill Profile
		err = rows.Scan(&id, &Profile.Uuid, &Profile.Email, &Profile.FirstName, &Profile.LastName, &Profile.FullName, &Profile.Mobile, &Profile.Birthdate, &Profile.Sex, &Profile.Updated, &Profile.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		// TODO ADD ROLES
		// get roles by id
		/*var roles []models.Role
		roles, err = GetRolesByProfileId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}

		// join roles to Profile
		Profile.Roles = roles
		// append to list of Profile*/

		// TODO ADD ADDRESSES
		// get roles by id
		var addresses []models.Address
		addresses, err = GetAddressesByProfileId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		Profile.Addresses = addresses

		// TODO ADD ENTITIES
		// get roles by id
		/*var roles []models.Role
		roles, err = GetRolesByProfileId(id)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}

		// join roles to Profile
		Profile.Roles = roles
		// append to list of Profile*/

		// join roles to Profile
		//Profile.Roles = roles
		// append to list of Profile*/

		Profiles = append(Profiles, *Profile)
	}
	return Profiles, err
}

/**
 * update Profile
 */
func UpdateProfile(Profile *models.Profile) (err error) {
	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//slect id
	rows, err := tx.Query("SELECT id FROM profile WHERE uuid = ?", Profile.Uuid)
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

	//update Profile
	_, err = tx.Exec("UPDATE profile SET updated = ? WHERE id = ?", time.Now().Unix(), id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	//update profile
	_, err = tx.Exec("UPDATE profile SET firstname = ?, lastname = ? WHERE id = ?", Profile.FirstName, Profile.LastName, id)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

/**
 * Create Profile
 */
func CreateProfile(Profile *models.ProfileCreate) (err error) {
	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}

	// Check for existing profile
	rows, err := tx.Query("SELECT id FROM profile WHERE email = ?", Profile.Email)
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
	uuid := uuid.New()
	_, err = tx.Exec("INSERT INTO profile (uuid, firstname, lastname, email, mobile, birthdate, sex, updated, created) VALUES "+
		"(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		uuid, Profile.FirstName, Profile.LastName, Profile.Email, Profile.Mobile, Profile.Birthdate, Profile.Sex, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

func DeleteProfile(deleteBody *models.DeleteBody) (err error) {
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//slect id
	rows, err := tx.Query("SELECT id FROM profile WHERE uuid = ?", deleteBody.Uuid)
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
	_, err = tx.Exec("DELETE FROM profile WHERE id = ?", id)
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
	rows, err := utils.DB.Query("SELECT id FROM profile WHERE uuid = ?", assign.Assign)
	if err != nil {
		log.Print("Database Error", err)
		return err
	}
	// select profile_id from rows
	var ProfileId int
	for rows.Next() {
		err = rows.Scan(&ProfileId)
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
	//select Profile_id from rows
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
	// insert Profile_has_Role
	_, err = tx.Exec("INSERT INTO Supporter_has_Role (Supporter_id, Role_Id) VALUES(?, ?)", ProfileId, roleId)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()
}
