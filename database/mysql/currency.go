package mysql

import (
	"database/sql"

	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/jmoiron/sqlx"
)
type Currency struct {
	cli *sqlx.DB
}

func (r *Currency) SelectByCurrencyCode(tx *sql.Tx, code model.CurrencyCode) (*model.Currency, bool, error) {
	var result model.Currency

	query := `
		SELECT
			currency_id,
			code,
			symbol,
			created_at,
			updated_at
		FROM um_help.tab_currency
		WHERE code = ?
		AND deleted_at IS NULL
		LIMIT 1;`

	exec := r.cli.Query
	if tx != nil {
		exec = tx.Query
	}

	rows, err := exec(query, code)
	if err != nil {
		return nil, false, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&result); err != nil {
			return nil, false, err
		}

	} else {
		return nil, false, nil
	}

	if rows.Err() != nil {
		return nil, false, rows.Err()
	}

	return &result, true, nil
}
