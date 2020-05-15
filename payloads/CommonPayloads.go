package payloads

import internalContracts "nikan.dev/pronto/internals/contracts"

type PaginationPayload struct {
	Page int
	PageSize int
}

func (payload PaginationPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.PositiveNumber(payload.Page, "Page"); err != nil {
		return err
	}
	if err := validator.PositiveNumber(payload.PageSize, "PageSize"); err != nil {
		return err
	}

	return  nil
}
type ChunkPayload struct {
	TotalRecord int
	TotalPage   int
	Records     interface{}
	Offset      int
	Limit       int
	Page        int
	PrevPage    int
	NextPage    int
}