package routes

import (
	"net/http"
	"github.com/JonasBrothers12/BackendChallenge/Validations"
	"github.com/gin-gonic/gin"
)

func HandleCreateAccount(c *gin.Context) {
	var Body validations.UserAccountRequest
	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.String(http.StatusBadRequest,err.Error())
		return
	}
	errname := validations.ValidateName(Body.FirstName,Body.LastName)
	if errname != nil {
		c.String(http.StatusBadRequest,errname.Error())
		return 
	}
	errtaxID := validations.ValidateTaxID(Body.TaxID)
	if errtaxID != nil {
		c.String(http.StatusBadRequest,errtaxID.Error())
		return 
	}
	c.String(http.StatusOK,"%s\n%s\n%s\n",Body.FirstName,Body.LastName,Body.TaxID)
}
func HandleAuth(c *gin.Context) {

}
func HandleTransactionRegister(c *gin.Context) {

}
func HandleTransactionReturn(c *gin.Context) {

}
func HandleViewTransaction(c *gin.Context) {

}
func HandleViewBalance(c *gin.Context) {

}

//Cadastro de conta
//Autenticação
//Cadastro de transação
//Estorno de transação
//Busca de transações por data
//Visualização de saldo
