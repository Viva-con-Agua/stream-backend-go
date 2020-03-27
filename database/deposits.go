package database

import (
    "../models"
    "../utils"
    _ "github.com/go-sql-driver/mysql"
    "github.com/google/uuid"
    "log"
    "time"
)

/**
 * select Deposit
 */
func GetDepositById(search string) (Deposits []models.Deposit, err error) {
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
        deposit := new(models.Deposit)

        //map row to model
        // TODO: Update model
        err = rows.Scan(&deposit.Uuid, &deposit.DateOfDeposit, &deposit.DateOfDeposit, &deposit.Created)
        if err != nil {
            log.Print("Database Error: ", err)
            return nil, err
        }
        Deposits = append(Deposits, *deposit)
    }
    return Deposits, err
}

/**
 * select list of Deposit
 */
func GetDepositList(page *models.Page, sort string, filter *models.FilterDeposit) (Deposits []models.Deposit, err error) {

    // execute the query
    DepositQuery := "SELECT Deposit.id, Deposit.uuid, Deposit.email, Deposit.firstname, Deposit.lastname, CONCAT(Deposit.firstname, ' ', Deposit.lastname) AS fullname, Deposit.mobile, Deposit.birthdate, Deposit.sex, Deposit.updated, Deposit.created " +
        "FROM Deposit LEFT JOIN Deposit_has_entity ON Deposit.id = Deposit_has_entity.Deposit_id " +
        "WHERE Deposit.email LIKE ? " +
        "GROUP BY Deposit.id " +
        sort + " " +
        "LIMIT ?, ?"
    rows, err := utils.DB.Query(DepositQuery, filter.Name, page.Offset, page.Count)
    if err != nil {
        log.Print("Database Error", err)
        return nil, err
    }

    // convert each row
    for rows.Next() {

        //create new Model
        deposit := new(models.Deposit)

        //map row to model
        // TODO: Update model
        err = rows.Scan(&deposit.Uuid, &deposit.Uuid, &deposit.Uuid, &deposit.Created)
        if err != nil {
            log.Print("Database Error: ", err)
            return nil, err
        }
        Deposits = append(Deposits, *deposit)
    }
    return Deposits, err
}

/**
 * select list of Deposit
 */
func GetDepositCount(page *models.Page, sort string, filter *models.FilterDeposit) (Deposits []models.Deposit, err error) {

    // execute the query
    DepositQuery := "SELECT Deposit.id, Deposit.uuid, Deposit.email, Deposit.firstname, Deposit.lastname, CONCAT(Deposit.firstname, ' ', Deposit.lastname) AS fullname, Deposit.mobile, Deposit.birthdate, Deposit.sex, Deposit.updated, Deposit.created " +
        "FROM Deposit LEFT JOIN Deposit_has_entity ON Deposit.id = Deposit_has_entity.Deposit_id " +
        "WHERE Deposit.email LIKE ? " +
        "GROUP BY Deposit.id " +
        sort + " " +
        "LIMIT ?, ?"
    rows, err := utils.DB.Query(DepositQuery, filter.Name, page.Offset, page.Count)
    if err != nil {
        log.Print("Database Error", err)
        return nil, err
    }

    // convert each row
    for rows.Next() {

        //create new Model
        deposit := new(models.Deposit)

        //map row to model
        // TODO: Update model
        err = rows.Scan(&deposit.Uuid, &deposit.Uuid, &deposit.Uuid, &deposit.Created)
        if err != nil {
            log.Print("Database Error: ", err)
            return nil, err
        }
        Deposits = append(Deposits, *deposit)
    }
    return Deposits, err
}

/**
 * Update Deposit
 */
func UpdateDeposit(Deposit *models.Deposit) (err error) {

    // sgl begin
    tx, err := utils.DB.Begin()
    if err != nil {
        log.Print("Database Error: ", err)
        return err
    }
    //update Deposit
    _, err = tx.Exec("UPDATE Deposit SET updated = ? WHERE uuid = ?", time.Now().Unix(), Deposit.Uuid)
    if err != nil {
        tx.Rollback()
        log.Print("Database Error: ", err)
        return err
    }
    return tx.Commit()

}

/**
 * Create Deposit
 */
func CreateDeposit(Deposit *models.DepositCreate) (err error) {

    // sgl begin
    tx, err := utils.DB.Begin()
    if err != nil {
        log.Print("Database Error: ", err)
        return err
    }

    // Insert Deposit
    uuid := uuid.New()
    _, err = tx.Exec("INSERT INTO Deposit (uuid, firstname, lastname, email, mobile, birthdate, sex, updated, created) VALUES "+
        "(?, ?, ?, ?, ?, ?, ?, ?, ?)",
        uuid, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, time.Now().Unix(), time.Now().Unix())
    if err != nil {
        tx.Rollback()
        log.Print("Database Error: ", err)
        return err
    }
    return tx.Commit()

}

/**
 * Create Deposit
 */
func ConfirmDeposit(Deposit *models.DepositConfirm) (err error) {

    // sgl begin
    tx, err := utils.DB.Begin()
    if err != nil {
        log.Print("Database Error: ", err)
        return err
    }

    // Insert Deposit
    uuid := uuid.New()
    _, err = tx.Exec("INSERT INTO Deposit (uuid, firstname, lastname, email, mobile, birthdate, sex, updated, created) VALUES "+
        "(?, ?, ?, ?, ?, ?, ?, ?, ?)",
        uuid, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, Deposit.DateOfDeposit, time.Now().Unix(), time.Now().Unix())
    if err != nil {
        tx.Rollback()
        log.Print("Database Error: ", err)
        return err
    }
    return tx.Commit()

}
