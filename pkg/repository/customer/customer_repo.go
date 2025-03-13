package customer_repo

type CustomerInterface interface {
	CheckDatabaseStats() map[string]string
}
