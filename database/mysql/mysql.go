package mysql

import (
	"fmt"
	"time"

	"github.com/JonasBrothers12/BackendChallenge/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func ConnectDatabase(cfg *config.MySQLConfig) (*sqlx.DB, error){

	url := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?paseTime=true"

	cli, err := sqlx.Open("mysql",url)

	if err != nil {
		fmt.Println(cli.DB.Stats())
		return nil,err
	}

	cli.DB.SetConnMaxLifetime(time.Minute * 5)
	cli.DB.SetMaxIdleConns(5)
	cli.DB.SetMaxOpenConns(100)
	
	if err := cli.Ping(); err != nil {
		fmt.Println(cli.DB.Stats())
		return nil,err
	}

	return cli,nil
}