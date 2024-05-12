package model

import "time"

type WalletUser struct {
	ID 			int64		`db:"wallet_id"`
	OwnerID		int64		`db:"owner_id"`			
	Alias		string		`db:"alias"`
	CurrencyID 	int64		`db:"currency_id"`
	Balance		int64		`db:"balance"`
	CreatedAt	time.Time	`db:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at"`
	DeletedAt	time.Time	`db:"deleted_at"`
}