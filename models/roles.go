package models

type Role uint

const (
	NullRole Role = iota
	Client
	Admin
)

func (userRole Role) String() string {

	names := [...]string{
		"NullRole",
		"Client",
		"Admin",
	}

	if userRole < NullRole || userRole > Admin {
		return "Unknown"
	}

	return names[userRole]
}
