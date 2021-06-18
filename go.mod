module wxauth

go 1.16

require (
	github.com/GeertJohan/go.rice v1.0.2
	github.com/go-redis/redis/v8 v8.9.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/mojocn/base64Captcha v1.3.4 // indirect
	golang.org/x/image v0.0.0-20210504121937-7319ad40d33e // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	wxauth/captcha v0.0.0
	wxauth/codegen v0.0.0 // indirect
	wxauth/datatype v0.0.0 // indirect
	wxauth/e2ee v0.0.0
	wxauth/handlers v0.0.0-00010101000000-000000000000
	wxauth/platform/dbmgr v0.0.0
	wxauth/redismgr v0.0.0 // indirect
)

replace (
	wxauth/captcha => ./captcha
	wxauth/codegen => ./codegen
	wxauth/datatype => ./datatype
	wxauth/e2ee => ./e2ee
	wxauth/encryptmgr => ./encryptmgr
	wxauth/handlers => ./handlers
	wxauth/mail => ./mail
	wxauth/mailmgr => ./mailmgr
	wxauth/platform/dbmgr => ./platform/dbmgr
	wxauth/redismgr => ./redismgr
)
