package response

type  UserAcount struct{
	ID        int64   `json:"userID"`
	Balance   int64	  `json:"balance"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	TaxID 	  string  `json:"taxID"`
	CreatedAt string  `json:"createdAt"`
	UpdateAt  string  `json:"updatAt"`
	DeleteAt  string  `json:"deleteAt"`
}
