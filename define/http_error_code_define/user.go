package http_error_code_define

// The authentication of the identity information failed,
// and the authorization could not be performed.
// When the authorization token or the renewal token is obtained,
// the error is reported when the identity fails.
const User_LoginIdentityInfoVerifyFailed = "user_login_identity_info_verify_failed"

// create user when user number already exists
const User_CreateFailedNumberAlreadyExists = "user_create_failed_number_already_exists"
