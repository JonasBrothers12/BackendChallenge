package user

import (
	"net/http"

	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database"
	"github.com/JonasBrothers12/BackendChallenge/presenter/requisition"
	"github.com/JonasBrothers12/BackendChallenge/validations"
	"github.com/gin-gonic/gin"
)

func HandleNewUser(c *gin.Context){
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
	DB,err := database.NewDatabaseAction(config.Getconfigs())
	if err != nil {
		c.String(http.StatusBadRequest,err.Error())		
		return
	}
	DB.MySQL.User.InsertNewUser(&Body)
}