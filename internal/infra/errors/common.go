package errors

const (
	InvalidIdError    CustomError = "Invalid id, must be integer above 0"
	InvalidParamError CustomError = "Invalid param error"
	DataNotFoundError CustomError = "Data does not exist"
)
