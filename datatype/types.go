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
	Email   string
	PassRSA string
}

type User struct {
	Email    string
	PassHash string
	UserRole Role
	Services []Service
}

type UserDataModel struct {
	ID       int
	Email    string
	PassHash string
	UserRole string
	Services string
}
