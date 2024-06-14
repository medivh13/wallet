package errors

// code for team internal error (custom error)
const (
	DATA_INVALID               ErrorCode = 1001
	INVALID_AMOUNT             ErrorCode = 1002
	STATUS_PAGE_NOT_FOUND      ErrorCode = 1003
	UNAUTHORIZED               ErrorCode = 1004
	USER_ALREADY_EXIST         ErrorCode = 1005
	DESTINATION_USER_NOT_FOUND ErrorCode = 1006
	INSUFICIENT_BALANCE        ErrorCode = 1007
	TO_OWN_ACCOUNT             ErrorCode = 1008

	// server error
	UNKNOWN_ERROR        ErrorCode = 2000
	FAILED_CREATE_DATA   ErrorCode = 2001
	FAILED_RETRIEVE_DATA ErrorCode = 2002
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	UNAUTHORIZED: {
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
	USER_ALREADY_EXIST: {
		ClientMessage: "Username Already Exist.",
		SystemMessage: "Username Already Exist.",
		ErrorCode:     USER_ALREADY_EXIST,
	},
	INVALID_AMOUNT: {
		ClientMessage: "Invalid Topup Amount",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     INVALID_AMOUNT,
	},
	DESTINATION_USER_NOT_FOUND: {
		ClientMessage: "Destination user not found",
		SystemMessage: "Destination user not found",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	INSUFICIENT_BALANCE: {
		ClientMessage: "Insufficient balance",
		SystemMessage: "Insufficient balance",
		ErrorCode:     INSUFICIENT_BALANCE,
	},
	TO_OWN_ACCOUNT: {
		ClientMessage: "Can't Transfer to own account",
		SystemMessage: "Can't Transfer to own account",
		ErrorCode:     TO_OWN_ACCOUNT,
	},
}
