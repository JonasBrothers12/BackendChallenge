package model

type  UserAcount struct{
	ID        int64   `db:"user_id"`
	Balance   int64	  `db:"balance"`
	FirstName string  `db:"first_name"`
	LastName  string  `db:"last_name"`
	TaxID 	  string  `db:"tax_id"`
	CreatedAt string  `db:"created_at"`
	UpdateAt  string  `db:"update_at"`
	DeleteAt  string  `db:"delete_at"`
}
