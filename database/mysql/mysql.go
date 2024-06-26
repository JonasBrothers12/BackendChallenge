package mysql

import (
	"fmt"
	"github.com/JonasBrothers12/BackendChallenge/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DataBase struct{
	User *User
	Cli *sqlx.DB
	Wallet *Wallet
	Currency *Currency
}

func ConnectDatabase(cfg *config.MySQLConfig) (*DataBase, error){

	url := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database

	cli, err := sqlx.Open("mysql",url)

	if err != nil {
		fmt.Println(cli.DB.Stats())
		return nil,err
	}
	if err := cli.Ping(); err != nil {
		fmt.Println(cli.DB.Stats())
		return nil,err
	}

	return &DataBase{
		User: &User{cli:cli},
		Cli: cli,
		Wallet: &Wallet{cli: cli},
		Currency: &Currency{cli:cli},
	}, nil
}