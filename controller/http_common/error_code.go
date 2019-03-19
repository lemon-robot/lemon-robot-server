package http_common

// Failed while parsing user uploaded file
const ErrCode_FileResource_AnalysisFailed = "file_resource_analysis_failed"

// The file_resource_key obtained from the header is invalid.
const ErrCode_FileResource_KeyInvalid = "file_resource_key_invalid"

// The authentication of the identity information failed,
// and the authorization could not be performed.
// When the authorization token or the renewal token is obtained,
// the error is reported when the identity fails.
const ErrCode_User_LoginIdentityInfoVerifyFailed = "user_login_identity_info_verify_failed"

// create user when user number already exists
const ErrCode_User_CreateFailedNumberAlreadyExists = "user_create_failed_number_already_exists"

// Common error, error inside the server, and do not want to expose the error details to the user
const ErrCode_Common_ServerInternalError = "common_server_internal_error"

// Common error, Return when the user accesses the interface that requires authorization and does not carry the legal authorization token.
const ErrCode_Common_Unauthorized = "common_unauthorized"
