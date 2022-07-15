package dao

import (
	"Pg_FW_Template/config"
	"Pg_FW_Template/model"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// TemplateDAO persists user data in database
type TemplateDAO struct{}

// NewUserDAO creates a new TemplateDAO
func NewUserDAO() *TemplateDAO {
	return &TemplateDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *TemplateDAO) Get(id uint) (*model.TemplateTest, error) {
	var templateTest model.TemplateTest

	var dbCfg = config.GetSourceCfg()
	dsn := dbCfg.GetDSN()
	log.Printf("DSN: %s\n", dsn)

	db, err := sql.Open(dbCfg.Driver, dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query("select * from template_test where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&templateTest.Id, &templateTest.RndStr, &templateTest.RndFloat, &templateTest.RndInt)
		if err != nil {
			return nil, err
		}
	}

	return &templateTest, err
}
