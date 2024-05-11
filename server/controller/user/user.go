package user

import (
	"fmt"
	"net/http"

	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database/mysql"
	"github.com/JonasBrothers12/BackendChallenge/presenter/requisition"
	"github.com/JonasBrothers12/BackendChallenge/validations"
	"github.com/gin-gonic/gin"
)

func HandleNewuser(c *gin.Context){
	var Body requisition.UserAccountRequest
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
	Db ,err := mysql.ConnectDatabase(config.Getconfigs().MySQLConfig)
	if err != nil {
		c.String(http.StatusBadRequest,err.Error())		
		return
	}
	fmt.Println("ok")
	err = mysql.InsertNewUser(Db,&Body)
	if err != nil {
		c.String(http.StatusBadRequest,err.Error())		
		return
	}
	
}