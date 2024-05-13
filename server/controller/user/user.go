package user

import (
	"net/http"
	/*
		"github.com/JonasBrothers12/BackendChallenge/config"
		"github.com/JonasBrothers12/BackendChallenge/database"
	*/
	"github.com/JonasBrothers12/BackendChallenge/presenter/requisition"
	"github.com/JonasBrothers12/BackendChallenge/service"
	"github.com/JonasBrothers12/BackendChallenge/util"
	"github.com/JonasBrothers12/BackendChallenge/validations"
	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	resutil *util.Util
	svc     *service.Service
}

func NewController(svc *service.Service, resutil *util.Util) *ControllerUser {
	return &ControllerUser{
		resutil: resutil,
		svc:     svc,
	}
}

func (ctrl *ControllerUser)HandleNewUser(c *gin.Context){
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
	ctrl.svc.User.CreateUserService(&Body)

}