package entities

type  UserAcount struct{
	ID        int64
	Balance   int64
	FirstName string
	LastName  string
	CPF 	  string	
	CreatedAt string
	UpdateAt  string
	DeleteAt  string
}

type UserTransaction struct{
	ID 		  				int64
	SenderID  				int64
	ReceiveID 				int64
	Amount_Transaction		int64
	TransactionAt		    string
}