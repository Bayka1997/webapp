package repository

import (
	"webapp/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountsRepositoryImpl struct {
	db *gorm.DB
}

const selectAccount = "select a.id as id, a.name as name, a.password as password," +
	" r.id as authority_id, r.name as authority_name " +
	" from account_master a inner join authority_master r on a.authority_id = r.id "

// Create implements AccountsRepository
func (*AccountsRepositoryImpl) Create(rep Repository) (*model.Account, error) {
	if err := rep.Select("name", "password", "authority_id").Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// FindByName implements AccountsRepository
func (*AccountsRepositoryImpl) FindByName(rep RepositoryImpl, name string) {
	var account *model.Account

	var rec RecordAccount
	rep.Raw(selectAccount+" where a.name = ?", name).Scan(&rec)
	account = convertToAccount(&rec)

	return account, nil
}

// NewAccount implements AccountsRepository
func (*AccountsRepositoryImpl) NewAccount(name string, password string, authorityID uint) {
	return &Account{Name: name, Password: password, AuthorityID: authorityID}
}

// NewAccountWithPlainPassword implements AccountsRepository
func (*AccountsRepositoryImpl) NewAccountWithPlainPassword(name string, password string, authorityID uint) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), config.PasswordHashCost)
	return &Account{Name: name, Password: string(hashed), AuthorityID: authorityID}
}

func NewAccountsREpositoryImpl(Db *gorm.DB) AccountsRepository {
	return &AccountsRepositoryImpl{db: Db}
}

// NewAccount is constructor.
func (req *AccountsRepositoryImpl) NewAccount(name string, password string, authorityID uint) *model.Account {
	return &model.Account{Name: name, Password: password, AuthorityID: authorityID}
}
