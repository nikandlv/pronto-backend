package internalContracts

type IConfiguration interface {
	Boot() error
	Get(key string) (interface{}, error)
	Contains(key string) bool
}
