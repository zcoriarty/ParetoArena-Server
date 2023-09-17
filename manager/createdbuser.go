package manager

import (
	"fmt"

	"github.com/zcoriarty/Backend/config"

	"github.com/go-pg/pg/v9"
)

// CreateDatabaseUserIfNotExist creates a database user
func CreateDatabaseUserIfNotExist(db *pg.DB, p *config.PostgresConfig) {
	statement := fmt.Sprintf(`SELECT * FROM pg_roles WHERE rolname = '%s';`, p.User)
	res, err := db.Exec(statement)
	if err != nil {
		fmt.Println("query execution error")
	}
	if res.RowsReturned() == 0 {
		statement = fmt.Sprintf(`CREATE USER %s WITH PASSWORD '%s';`, p.User, p.Password)
		_, err := db.Exec(statement)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf(`Created user %s`, p.User)
		}
	} else {
		fmt.Printf("Database user %s already exists\n", p.User)
	}
}
