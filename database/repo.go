package database

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database/mysql"
)

type DatabaseManager struct{
	MySQL *mysql.DataBase
}

func NewDatabaseAction(cfg *config.Config) (*DatabaseManager,error){
	mysql, err := mysql.ConnectDatabase(cfg.MySQLConfig)
	if err != nil {
		return nil, err
	}

	return &DatabaseManager{
		MySQL: mysql,
	},nil

}