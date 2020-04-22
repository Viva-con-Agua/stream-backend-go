package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"stream-backend-go/models"
	"stream-backend-go/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
		err = rows.Scan(&taking.Uuid, &taking.Comment, &taking.Author)
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
func GetTakingList(page *models.Page, sort string, filter *models.FilterTaking) (takings []models.Taking, err error) {

	// execute the query
	TakingQuery := "SELECT DISTINCT t.id, t.uuid, t.received, t.description, t.comment, t.category, " +
		"JSON_OBJECT('uuid', a.uuid, 'name', a.name), " +
		"CONCAT('[', GROUP_CONCAT(JSON_OBJECT('uuid', i.uuid, 'name', i.name)), ']'), " +
		"CONCAT('[', GROUP_CONCAT(JSON_OBJECT('uuid', s.uuid, 'category', s.category, 'description', s.description, " +
		"'amount', s.amount, 'currency', s.currency, 'typeOfSource', s.type_of_source, 'norms', s.norms)), ']'), " +
		"CONCAT('[', GROUP_CONCAT(JSON_OBJECT('amount', m.amount, 'currency', m.currency, " +
		"'status', m.status, 'updated', m.updated, 'created', m.created)), ']'), " +
		"CONCAT('[', GROUP_CONCAT(JSON_OBJECT('uuid', c.uuid, 'name', c.name)), ']') " +
		"FROM taking AS t " +
		"JOIN user_taking AS a ON t.id = a.taking_id AND a.tag = 'author' " +
		"JOIN user_taking AS i ON t.id = i.taking_id " +
		"JOIN source AS s ON t.id = s.taking_id " +
		"JOIN money AS m ON t.id = m.taking_id " +
		"JOIN crew_taking AS c ON t.id = c.taking_id " +
		"WHERE t.description like '%' " +
		"GROUP BY t.id, a.uuid, a.name " +
		sort + " " +
		"LIMIT ?, ?"
	rows, err := utils.DB.Query(TakingQuery, page.Offset, page.Count)
	if err != nil {
		log.Print("Database Error", err)
		return nil, err
	}

	// convert each row
	for rows.Next() {

		//create new Model
		taking := new(models.Taking)
		var id int
		var authorBytes []byte
		var supporterBytes []byte
		var sourceBytes []byte
		var moneyBytes []byte
		var crewBytes []byte

		//map row to model
		// TODO: Update model
		err = rows.Scan(&id, &taking.Uuid, &taking.Received, &taking.Description,
			&taking.Comment, &taking.Category, &authorBytes, &supporterBytes, &sourceBytes, &moneyBytes, &crewBytes)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		//convert author
		author := new(models.User)
		err = json.Unmarshal(authorBytes, &author)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		taking.Author = *author

		//convert supporter
		supporter := new(models.UserList)
		err = json.Unmarshal(supporterBytes, &supporter)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		taking.User = *supporter.Distinct()

		//convert sources
		source := new(models.SourceList)

		err = json.Unmarshal(sourceBytes, &source)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		taking.Sources = *source.Distinct()

		//convert money
		money := new(models.MoneyList)
		err = json.Unmarshal(moneyBytes, &money)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		taking.Money = *money.Distinct()

		//convert crew
		crew := new(models.CrewList)
		err = json.Unmarshal(crewBytes, &crew)
		if err != nil {
			log.Print("Database Error: ", err)
			return nil, err
		}
		taking.Crews = *crew.Distinct()

		takings = append(takings, *taking)
	}
	return takings, err
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

func GetOrInsertUser(u *models.User, tx *sql.Tx) (id int64, err error) {
	rows, err := tx.Query("SELECT id FROM user WHERE user.uuid = ? LIMIT 1", u.Uuid)
	if err != nil {
		tx.Rollback()
		log.Print("GetOrInsertUser.SelectUser: ", err)
	}
	for rows.Next() {
		err = rows.Scan(&id)
	}
	if id == 0 {
		res, err := tx.Exec("INSERT INTO user (uuid, name, updated, created) VALUES (?, ?, ?, ?)",
			u.Uuid, u.Name, time.Now().Unix(), time.Now().Unix())
		if err != nil {
			tx.Rollback()
			log.Print("GetOrInsertUser.InsertUser: ", err)
		}
		return res.LastInsertId()
	} else {
		return id, err
	}
}

func GetOrInsertCrew(c *models.Crew, tx *sql.Tx) (id int64, err error) {
	rows, err := tx.Query("SELECT id FROM crew WHERE crew.uuid = ? LIMIT 1", c.Uuid)
	if err != nil {
		tx.Rollback()
		log.Print("Database Error: ", err)
	}
	for rows.Next() {
		err = rows.Scan(&id)
	}
	if id == 0 {
		res, err := tx.Exec("INSERT INTO crew (uuid, name, updated, created) VALUES (?, ?, ?, ?)",
			c.Uuid, c.Name, time.Now().Unix(), time.Now().Unix())
		if err != nil {
			tx.Rollback()
			log.Print("Database Error: ", err)
		}
		return res.LastInsertId()
	} else {
		return id, err
	}
}

/**
 * Create Taking
 */
func CreateTaking(taking *models.TakingCreate) (err error) {

	// sgl begin
	tx, err := utils.DB.Begin()
	if err != nil {
		log.Print("CreateTaking.BeginDatabase: ", err)
		return err
	}
	// Insert Taking
	t_uuid := uuid.New()
	res, err := tx.Exec("INSERT INTO taking (uuid, received, description, category, reason_for_payment) VALUES "+
		"(?, ?, ?, ?, ?)",
		t_uuid, taking.Received, taking.Description, taking.Category, taking.ReasonForPayment)
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.InsertTaking: ", err)
		return err
	}

	// get taking_id
	taking_id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.getTakingId: ", err)
		return err
	}

	//insert author
	author_id, err := GetOrInsertUser(&taking.Author, tx)
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.getAuthorId: ", err)
		return err
	}
	_, err = tx.Exec("INSERT INTO taking_has_user (taking_id, user_id, tag) VALUES (?, ?, ?)", taking_id, author_id, "AUTHOR")
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.insertAuthor: ", err)
		return err
	}

	//insert involved User
	var user_id int64
	for _, s := range taking.User {
		user_id, err = GetOrInsertUser(&s, tx)
		if err != nil {
			tx.Rollback()
			log.Print("CreateTaking.getUserId: ", err)
			return err
		}
		_, err = tx.Exec("INSERT INTO taking_has_user (taking_id, user_id, tag) VALUES (?, ?, ?)", taking_id, user_id, "INVOLVED")
		if err != nil {
			tx.Rollback()
			log.Print("CreateTaking.InsertUser: ", err)
			return err
		}
	}
	//insert involved Crew
	var crew_id int64
	for _, s := range taking.Crews {
		crew_id, err = GetOrInsertCrew(&s, tx)
		if err != nil {
			tx.Rollback()
			log.Print("CreateTaking.GetCrewId: ", err)
			return err
		}
		_, err = tx.Exec("INSERT INTO taking_has_crew (taking_id, crew_id) VALUES (?, ?)", taking_id, crew_id)
		if err != nil {
			tx.Rollback()
			log.Print("CreateTaking.InsertCrew: ", err)
			return err
		}
	}
	// insert Sources
	for _, s := range taking.Sources {
		s_uuid := uuid.New()
		_, err = tx.Exec("INSERT INTO source (uuid, amount, currency, description, norms, category, type_of_source, taking_id) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			s_uuid, s.Amount, s.Currency, s.Description, s.Norms, s.Category, s.TypeOfSource.Category, taking_id)
		if err != nil {
			tx.Rollback()
			log.Print("CreateTaking.InsertSources: ", err)
			return err
		}
	}
	// create amount
	amount := 0
	currency := ""
	for _, v := range taking.Sources {
		amount = amount + v.Amount
		currency = v.Currency
	}
	_, err = tx.Exec("INSERT INTO money (amount, currency, status, updated, created, taking_id) "+
		"VALUES (?, ?, ?, ?, ?, ?)",
		amount, currency, "OPEN", time.Now().Unix(), time.Now().Unix(), taking_id)
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.InsertMoney: ", err)
		return err
	}

	//insert comment
	comment_id, err := InsertCommentWithTx(&taking.Comment, tx)
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.getCommentId: ", err)
		return err
	}
	_, err = tx.Exec("INSERT INTO taking_has_comment (taking_id, comment_id) VALUES (?, ?)", taking_id, comment_id)
	if err != nil {
		tx.Rollback()
		log.Print("CreateTaking.insertTakingsHasComment: ", err)
		return err
	}

	return tx.Commit()
}
