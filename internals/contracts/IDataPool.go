package internalContracts

type IDataPool interface {
	Boot(config IConfiguration, models ...interface{}) interface{}
}
