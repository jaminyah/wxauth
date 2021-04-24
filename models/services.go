package models

type Service uint

const (
	NullSvc Service = iota
	Notify
	WebRTC
	UserLogs
)

func (service Service) String() string {

	names := [...]string{
		"Unknown",
		"Notify",
		"WebRTC",
		"UserLogs",
	}

	if service < NullSvc || service > UserLogs {
		return "Unknown"
	}

	return names[service]
}
