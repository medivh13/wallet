package errors

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : wallet
 */

const (
	UNKNOWN_ERROR              ErrorCode = 0
	DATA_INVALID               ErrorCode = 4001
	FAILED_RETRIEVE_DATA       ErrorCode = 4002
	STATUS_PAGE_NOT_FOUND      ErrorCode = 4003
	INVALID_HEADER_X_API_KEY   ErrorCode = 4004
	UNAUTHORIZED               ErrorCode = 4005
	FAILED_CREATE_DATA         ErrorCode = 4006
	USER_ALREADY_EXIST         ErrorCode = 4007
	INVALID_AMOUNT             ErrorCode = 4008
	DESTINATION_USER_NOT_FOUND ErrorCode = 4009
	INSUFICIENT_BALANCE        ErrorCode = 4010
	TO_OWN_ACCOUNT             ErrorCode = 4011
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
