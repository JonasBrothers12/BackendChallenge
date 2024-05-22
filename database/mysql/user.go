package mysql

import (
	"database/sql"
	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/jmoiron/sqlx"
)

type User struct{
	cli *sqlx.DB
}

func (u *User) InsertNewUser(user *model.UserAcount,tx *sql.Tx) (int64,error){
	exec := u.cli.Exec
	if tx != nil{
		exec = tx.Exec
	}
	valueId,err := exec("INSERT INTO distopia_example.tab_user (first_name, last_name, tax_id) VALUES(?, ?, ?)",user.FirstName, user.LastName, user.TaxID)
	if err != nil {
		return 0,err
	}
	ValueID,err := valueId.LastInsertId()
	if err != nil {
		return 0,err
	}
	return ValueID,nil
}


