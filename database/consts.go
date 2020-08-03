package database

// dynamodb field names
const (
	keyNameRedirectsTable = "redirect_from"
	attrNameDescription   = "description"
	attrNameRedirect      = "redirect_to"
	attrNamePromotable    = "promotable"
	attrNameCode          = "code"
	attrNameCreated       = "created"
	attrNameModified      = "modified"
	attrNamePathMap       = "cpath_map"
	keyNameCertCacheTable = "domain"
	attrNameCacheValue    = "cert"
	keyNameUsersTable     = "username"
	attrNamePwd           = "pwd"
)

// default values
const (
	defaultDynamoUser     = "admin"
	defaultDynamoPassword = "$2a$12$gh.TtSizoP0JFLHACOdIouPr42713m6k/8fH8jKPl0xQAUBk0OIdS"
)
