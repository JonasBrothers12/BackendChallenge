package mysql

import (
	"database/sql"

	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct{
	cli *sqlx.DB
}

func (w *Wallet) InsertNewWallet(wallet *model.WalletUser,tx *sql.Tx) error{
	exec := w.cli.Exec
	if tx != nil{
		exec = tx.Exec
	}

	_, err := exec("INSERT INTO distopia_example.tab_wallet (owner_id,alias,currency_id,balance) VALUES(?, ?, ?, ?)",wallet.OwnerID,wallet.Alias,wallet.CurrencyID,wallet.Balance)
	
	if err != nil {
		return err
	}
	return nil
}