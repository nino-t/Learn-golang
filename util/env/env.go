package env

import "os"

const (
	development = "development"
	alpha       = "alpha"
	staging     = "staging"
	production  = "production"

	idevelopment int8 = iota + 1
	ialpha
	istaging
	iproduction
)

var (
	e  string
	ie int8
)

func init() {
	initEnv()
}

func initEnv() {
	e = os.Getenv("APP_ENV")
	switch e {
	case production:
		ie = iproduction
	case staging:
		ie = istaging
	case alpha:
		ie = ialpha
	default:
		e = development
		ie = idevelopment
	}
}

func Info() string {
	return e
}

func IsDevelopment() bool {
	return ie == idevelopment
}

func IsAlpha() bool {
	return ie == ialpha
}

func IsStaging() bool {
	return ie == istaging
}

func IsProduction() bool {
	return ie == iproduction
}
