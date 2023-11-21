package migrator

import (
	"database/sql"
	"fmt"
	"gameapp/repository/mysql"
	"github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect   string
	dbConfig  mysql.Config
	migration *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}

	return Migrator{dbConfig: dbConfig, dialect: "mysql", migration: migrations}
}
func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db:%v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migration, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migration:%v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db:%v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migration, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migration:%v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

//func (m Migrator) Status() {
//	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
//		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
//	if err != nil {
//		panic(fmt.Errorf("can't open mysql db:%v", err))
//	}
//
//	n, err := migrate.Exec(db, "mysql", m.migration, migrate.Up)
//	if err != nil {
//		panic(fmt.Errorf("can't apply migration:%v", err))
//	}
//	fmt.Printf("Applied %d migrations!\n", n)
//}
