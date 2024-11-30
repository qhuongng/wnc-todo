package http_common

type errorResponseCode struct {
	InvalidRequest      string
	InternalServerError string
	RecordNotFound      string
	MissingParameter    string
	InvalidDataType     string
	Unauthorized        string
}

var ErrorResponseCode = errorResponseCode{
	InvalidRequest:      "INVALID_REQUEST",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	RecordNotFound:      "RECORD_NOT_FOUND",
	MissingParameter:    "MISSING_PARAMETER",
	InvalidDataType:     "INVALID_DATA_TYPE",
	Unauthorized:        "UNAUTHORIZED",
}

type customValidationErrCode map[string]string

var CustomValidationErrCode = customValidationErrCode{
	"userrequest.role":  "INVALID_USER_ROLE",
	"userrequest.phone": "INVALID_PHONE_NUMBER",
}

type errorMessage struct {
	GormRecordNotFound string
	InvalidDataType    string
	InvalidRequest     string
}

var ErrorMessage = errorMessage{
	GormRecordNotFound: "record not found",
	InvalidDataType:    "invalid data type",
	InvalidRequest:     "invalid request",
}
