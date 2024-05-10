package routes

import (
	"database/sql"

	"net/http"

	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/JonasBrothers12/BackendChallenge/presenter/requisition"
	validations "github.com/JonasBrothers12/BackendChallenge/validations"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func HandleCreateAccount(c *gin.Context) {
	var Body other.UserAccountRequest
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
	user := &model.UserAcount{
		FirstName: Body.FirstName,
		LastName:  Body.LastName,
		TaxID: 	   Body.TaxID,
		Balance:   0,
	}
	cfg := config.Get()
	url := cfg.MySQLConfig.Username + ":" + cfg.MySQLConfig.Password + "@tcp(" + cfg.MySQLConfig.Host + ":" + cfg.MySQLConfig.Port + ")/" + cfg.MySQLConfig.Database + "?parseTime=true"
	db, err := sql.Open("mysql", url)
	if err != nil {
        panic(err.Error())
    }
	stmt, err := db.Prepare("INSERT INTO pupper_example.tab_user (first_name, last_name, document_number, balance) VALUES(?, ?, ?, ?)")
	if err != nil {
	    panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.TaxID, user.Balance)
	if err != nil {
		panic(err.Error())
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
