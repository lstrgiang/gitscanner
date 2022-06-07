package errors

const (
	InvalidIdError    CustomError = "Invalid id, must be integer above 0"
	InvalidParamError CustomError = "Invalid param error"
	InvalidUrlError   CustomError = "Invalid Url format"
	DataNotFoundError CustomError = "Data does not exist"
)
