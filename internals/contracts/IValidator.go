package internalContracts

type IValidator interface {
	ShortText(string, field string) error
	Text(string, field string) error
	ID(uint uint, field string) error
	PositiveNumber(num int, field string) error
	Timestamp(int64 int64, field string) error
	Email(string, field string) error
}
