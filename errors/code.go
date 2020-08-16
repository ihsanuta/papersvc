package errors

const (
	CodeValue       = 100
	CodeSQL         = 200
	CodeHTTPClient  = 500
	CodeHTTPHandler = 800
)

const (
	// Error On Values
	CodeValueInvalid = Code(iota + CodeValue)

	// Error On SQL
	CodeSQLBuilder = Code(iota + CodeSQL)
	CodeSQLRead
	CodeSQLRowScan
	CodeSQLCreate
	CodeSQLUpdate
	CodeSQLDelete
	CodeSQLUnlink
	CodeSQLTxBegin
	CodeSQLTxCommit
	CodeSQLPrepareStmt
	CodeSQLRecordMustExist
	CodeSQLCannotRetrieveLastInsertID
	CodeSQLCannotRetrieveAffectedRows
	CodeSQLUniqueConstraint
	CodeSQLRecordDoesNotMatch
	CodeSQLRecordIsExpired

	// Error on HTTP Client
	CodeHTTPClientMarshal = Code(iota + CodeHTTPClient)
	CodeHTTPClientUnmarshal
	CodeHTTPClientErrorOnRequest
	CodeHTTPClientErrorOnReadBody

	// Code HTTP Handler
	CodeHTTPBadRequest = Code(iota + CodeHTTPHandler)
	CodeHTTPNotFound
	CodeHTTPUnauthorized
	CodeHTTPInternalServerError
	CodeHTTPUnmarshal
	CodeHTTPMarshal
)
