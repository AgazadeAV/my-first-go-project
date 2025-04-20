package util

const (
	NameMinLength     = 1
	NameMaxLength     = 50
	UserNameMinLength = 1
	UserNameMaxLength = 50
	EmailMinLength    = 6
	EmailMaxLength    = 100
	PhoneNumberLength = 12
	PhonePattern      = `^\+7\d{10}$`
	NamePattern       = `^[A-Za-z]([A-Za-z' -]*[A-Za-z'])?$`
	UsernamePattern   = `^[A-Za-z0-9_-]+$`
	EmailPattern      = `^[A-Za-z0-9._-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`
	MinimalBirthYear  = 1900
	MinimalUserAge    = 18
)
