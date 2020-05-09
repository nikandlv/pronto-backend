package internalContracts

type IValidator interface {
	ShortText(string) error
	Text(string) error
	ID(uint) error
	Timestamp(int64) error
	Email(string) error
}
