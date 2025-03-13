package customer_repo

import "github.com/uptrace/bun"

type CustomerRepo struct {
	db *bun.DB
}

func NewCustomerRepo(db *bun.DB) CustomerInterface {
	return &CustomerRepo{
		db: db,
	}
}

func (r *CustomerRepo) CheckDatabaseStats() map[string]string {
	return nil
}
