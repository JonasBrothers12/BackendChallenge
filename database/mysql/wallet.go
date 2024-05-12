package mysql

import (
	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/JonasBrothers12/BackendChallenge/presenter/requisition"
	"github.com/jmoiron/sqlx"
)

func NewWalletUser(db *sqlx.DB,req *requisition.UserAccountRequest) error{
	user := &model.UserAcount{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		TaxID:     req.TaxID,
	}
	stmt, err := db.Prepare("INSERT INTO distopia_example.tab_wallet (first_name, last_name, tax_id) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.TaxID)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
