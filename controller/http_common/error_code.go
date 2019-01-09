package http_common

// Failed while parsing user uploaded file
const ErrCode_FileResource_AnalysisFailed = "file_resource_analysis_failed"

// The file_resource_key obtained from the header is invalid.
const ErrCode_FileResource_KeyInvalid = "file_resource_key_invalid"

// Common error, error inside the server, and do not want to expose the error details to the user
const ErrCode_Common_ServerInternalError = "common_server_internal_error"

// Common error, Return when the user accesses the interface that requires authorization and does not carry the legal authorization token.
const ErrCode_Common_Unauthorized = "common_unauthorized"
