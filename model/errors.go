package model

// Error string contants
const (
	ErrBodyUnmarshal       = "Body could not be unmarshaled"
	ErrRedirectMarshal     = "Redirect entry could not be marshaled"
	ErrRedirectsMarshal    = "Redirect entries could not be marshaled"
	ErrRedirectCompress    = "Redirect entry could not be compressed"
	ErrRedirectsDecompress = "Redirect entries could not be decompressed"
	ErrNilPointer          = "Unexpected nil pointer"
)

// Error string constants - Gzip
const (
	ErrWriterCreate = "Gzip writer could not be created"
	ErrWriterFlush  = "Gzip writer could not be flushed"
	ErrWriterClose  = "Gzip writer could not be closed"
	ErrReaderCreate = "Gzip reader could not be created"
	ErrReaderRead   = "Gzip reader could not be read"
	ErrReaderClose  = "Gzip reader could not be closed"
)

// Error string constants - Path maps
const (
	ErrPathMapsUnmarshal = "Path maps could not be unmarshaled"
	ErrPathMapsMarshal   = "Path maps could not be marshaled"
	ErrPathMapsEmpty     = "Path maps are empty"
)

// Error string constants - Auth
const (
	ErrJWTSign    = "JWT Token could not be signed"
	ErrPwdInvalid = "The password is invalid"
)

// Error string constants - Validator
const (
	ErrInvalidDomain   = "Invalid domain name"
	ErrInvalidTarget   = "Invalid redirect target"
	ErrInvalidHTTPCode = "Invalid HTTP code"
	ErrInvalidMeta     = "Invalid meta - modification/creation time"
)
