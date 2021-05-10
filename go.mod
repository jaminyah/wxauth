module wxauth

go 1.15

require (
	github.com/GeertJohan/go.rice v1.0.2
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/mojocn/base64Captcha v1.3.4
	github.com/stretchr/testify v1.7.0

	wxauth/handlers v0.0.0
)

replace (
	wxauth/handlers => ./handlers
)
