package datatype

type User struct {
	Email    string
	Password string
	UserRole Role
	Services []Service
}
