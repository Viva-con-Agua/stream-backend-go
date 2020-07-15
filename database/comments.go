package database

import (
	"database/sql"
	"log"
	"stream-backend-go/models"
	"time"

	"github.com/google/uuid"
)

// insert comment with given tx
func InsertCommentWithTx(c *models.CommentCreate, tx *sql.Tx) (id int64, err error) {
	var user_id int64

	//get user
	rows, err := tx.Query("SELECT id FROM user WHERE user.uuid = ? LIMIT 1", c.User.Uuid)
	if err != nil {
		tx.Rollback()
		log.Print("InsertCommentWithTx.selectUser: ", err)
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&user_id)
	}
	if user_id == 0 {
		res, err := tx.Exec("INSERT INTO user (uuid, name, updated, created) VALUES (?, ?, ?, ?)",
			c.User.Uuid, c.User.Name, time.Now().Unix(), time.Now().Unix())
		if err != nil {
			tx.Rollback()
			log.Print("InsertCommentWithTx.insertUser: ", err)
			return 0, err
		}
		user_id, err = res.LastInsertId()
	}
	c_uuid := uuid.New()
	res, err := tx.Exec("INSERT INTO comment (uuid, text, tag, created, user_id) VALUES (?, ?, ?, ?, ?)",
		c_uuid, c.Text, c.Tag, time.Now().Unix(), user_id)
	if err != nil {
		tx.Rollback()
		log.Print("InsertCommentWithTx.insertComment: ", err)
		return 0, err
	}
	return res.LastInsertId()

}
