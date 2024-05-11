package requisition

type UserAccountRequest struct{
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	TaxID 	  string `json:"taxID"`
}