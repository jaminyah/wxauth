package datatype

import (
	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type ConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type ActivateForm struct {
	Email string
	Code  string
}

type RegisterForm struct {
	Email    string
	Password string
}

type LoginForm struct {
	Email    string
	Password string
}

type User struct {
	Email    string
	Password string
	UserRole Role
	Services []Service
}

type UserDataModel struct {
	ID       int
	Email    string
	PassRSA  string
	UserRole string
	Services string
}
