module wxauth

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.7
	wxauth/captcha v0.0.0
	wxauth/codegen v0.0.0 // indirect
	wxauth/datatype v0.0.0 // indirect
	wxauth/handlers v0.0.0-00010101000000-000000000000
	wxauth/mailmgr v0.0.0 // indirect
	wxauth/platform/database v0.0.0
)

replace (
	wxauth/captcha => ./captcha
	wxauth/codegen => ./codegen
	wxauth/datatype => ./datatype
	wxauth/handlers => ./handlers
	wxauth/mail => ./mail
	wxauth/mailmgr => ./mailmgr
	wxauth/platform/database => ./platform/database
)
