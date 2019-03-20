package http_error_code_define

// Common error, error inside the server, and do not want to expose the error details to the user
const Common_ServerInternalError = "common_server_internal_error"

// Common error, Return when the user accesses the interface that requires authorization and does not carry the legal authorization token.
const Common_Unauthorized = "common_unauthorized"
