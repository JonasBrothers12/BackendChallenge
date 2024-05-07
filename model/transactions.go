package model

type UserTransaction struct{
	ID 		  				int64 `db:"user_id"`
	SenderID  				int64 `db:"sender_id"`
	ReceiveID 				int64 `db:"receive_id"`
	Amount_Transaction		int64 `db:"amount_transaction"`
	TransactionAt		    string `db:"transaction_at"`
}