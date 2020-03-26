package database

import (
	"stream-backend-go/models"
	"stream-backend-go/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"time"
)

/**
 * select Taking
 */
func GetTakingById(search string) (Takings []models.Taking, err error) {
	// Execute the Query
	query := "SELECT r.uuid, r.name, r.service_name, r.created " +
		"FROM Role AS r " +
		"WHERE r.uuid = ? "
	rows, err := utils.DB.Query(query, search)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}

	// convert each row
	for rows.Next() {

		//create new Model
		taking := new(models.Taking)

		//map row to model
		// TODO: Update model
		err = rows.Scan(&taking.Uuid, &taking.Comment, &taking.Author, &taking.Updated, &taking.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		Takings = append(Takings, *taking)
	}
	return Takings, err
}

/**
 * select list of Taking
 */
func GetTakingList(page *models.Page, sort string, filter *models.FilterTaking) (Takings []models.Taking, err error) {

	// execute the query
	TakingQuery := "SELECT Taking.id, Taking.uuid, Taking.email, Taking.firstname, Taking.lastname, CONCAT(Taking.firstname, ' ', Taking.lastname) AS fullname, Taking.mobile, Taking.birthdate, Taking.sex, Taking.updated, Taking.created " +
		"FROM Taking LEFT JOIN Taking_has_entity ON Taking.id = Taking_has_entity.Taking_id " +
		"WHERE Taking.email LIKE ? " +
		"GROUP BY Taking.id " +
		sort + " " +
		"LIMIT ?, ?"
	rows, err := utils.DB.Query(TakingQuery, filter.Name, page.Offset, page.Count)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}

	// convert each row
	for rows.Next() {

		//create new Model
		taking := new(models.Taking)

		//map row to model
		// TODO: Update model
		err = rows.Scan(&taking.Uuid, &taking.Comment, &taking.Author, &taking.Updated, &taking.Created)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		Takings = append(Takings, *taking)
	}
	return Takings, err
}

/**
 * Update Taking
 */
func UpdateTaking(Taking *models.Taking) (err error) {

	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}
	//update taking
	_, err = tx.Exec("UPDATE Taking SET updated = ? WHERE uuid = ?", time.Now().Unix(), Taking.Uuid)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}

/**
 * Create Taking
 */
func CreateTaking(Taking *models.TakingCreate) (err error) {

	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("Database Error: ", err)
		return err
	}

	// Insert Taking
	uuid := uuid.New()
	_, err = tx.Exec("INSERT INTO Taking (uuid, firstname, lastname, email, mobile, birthdate, sex, updated, created) VALUES "+
		"(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		uuid, Taking.Email, Taking.Email, Taking.Email, Taking.Email, Taking.Email, Taking.Email, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
		return err
	}
	return tx.Commit()

}
