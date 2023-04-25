package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewREpositoryImpl(Db *gorm.DB) Repository {
	return &RepositoryImpl{db: Db}
}

// Select specify fields that you want to retrieve from database when querying, by default, will select all fields;
func (rep *RepositoryImpl) Select(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Select(query, args...)
}

// Find find records that match given conditions.
func (rep *RepositoryImpl) Find(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.Find(out, where...)
}

// Exec exec given SQL using by gorm.DB.
func (rep *RepositoryImpl) Exec(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Exec(sql, values...)
}

// First returns first record that match given conditions, order by primary key.
func (rep *RepositoryImpl) First(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.First(out, where...)
}

// Raw returns the record that executed the given SQL using gorm.DB.
func (rep *RepositoryImpl) Raw(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Raw(sql, values...)
}

// Create insert the value into database.
func (rep *RepositoryImpl) Create(value interface{}) *gorm.DB {
	return rep.db.Create(value)
}

// Save update value in database, if the value doesn't have primary key, will insert it.
func (rep *RepositoryImpl) Save(value interface{}) *gorm.DB {
	return rep.db.Save(value)
}

// Update update value in database
func (rep *RepositoryImpl) Updates(value interface{}) *gorm.DB {
	return rep.db.Updates(value)
}

// Delete delete value match given conditions.
func (rep *RepositoryImpl) Delete(value interface{}) *gorm.DB {
	return rep.db.Delete(value)
}

// Where returns a new relation.
func (rep *RepositoryImpl) Where(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Where(query, args...)
}

// Preload preload associations with given conditions.
func (rep *RepositoryImpl) Preload(column string, conditions ...interface{}) *gorm.DB {
	return rep.db.Preload(column, conditions...)
}

// Scopes pass current database connection to arguments `func(*DB) *DB`, which could be used to add conditions dynamically
func (rep *RepositoryImpl) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return rep.db.Scopes(funcs...)
}

// ScanRows scan `*sql.Rows` to give struct
func (rep *RepositoryImpl) ScanRows(rows *sql.Rows, result interface{}) error {
	return rep.db.ScanRows(rows, result)
}

// Close close current db connection. If database connection is not an io.Closer, returns an error.
func (rep *RepositoryImpl) Close() error {
	sqlDB, _ := rep.db.DB()
	return sqlDB.Close()
}

// DropTableIfExists drop table if it is exist
func (rep *RepositoryImpl) DropTableIfExists(value interface{}) error {
	return rep.db.Migrator().DropTable(value)
}

// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
func (rep *RepositoryImpl) AutoMigrate(value interface{}) error {
	return rep.db.AutoMigrate(value)
}

// Model specify the model you would like to run db operations
func (rep *RepositoryImpl) Model(value interface{}) *gorm.DB {
	return rep.db.Model(value)
}

// Transaction start a transaction as a block.
func (rep *RepositoryImpl) Transaction(fc func(tx Repository) error) (err error) {
	panicked := true
	tx := rep.db.Begin()
	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txrep := &RepositoryImpl{}
	txrep.db = tx
	err = fc(txrep)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
