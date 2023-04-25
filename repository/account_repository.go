package repository

import "webapp/model"

type RecordAccount struct {
	ID            uint
	Name          string
	Password      string
	AuthorityID   uint
	AuthorityName string
}

type AccountsRepository interface {
	NewAccount(name string, password string, authorityID uint)
	NewAccountWithPlainPassword(name string, password string, authorityID uint)
	FindByName(rep RepositoryImpl, name string)
	Create(rep Repository) (*model.Account, error)
}
