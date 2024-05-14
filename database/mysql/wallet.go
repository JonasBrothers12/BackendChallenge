package mysql

import (
	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct{
	cli *sqlx.DB
}

func (w *Wallet) InsertNewWallet(wallet *model.WalletUser) error{
	stmt, err := w.cli.Prepare("INSERT INTO distopia_example.tab_wallet (owner_id,alias,currency_id,balance) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(wallet.OwnerID,wallet.Alias,wallet.CurrencyID,wallet.Balance)
	if err != nil {
		panic(err.Error())
	}
	return nil
}