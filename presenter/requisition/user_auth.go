package requisition

type AuthUser struct{
	DocumentNumber string `json:"documentNumber"`
	Password       string `json:"password"`
}