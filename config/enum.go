package config

type UserRole string

var EnumUserRole = struct {
	USER  UserRole
	ADMIN UserRole
}{
	USER:  "user",
	ADMIN: "admin",
}

type NotificationType string

var EnumNotificationType = struct {
	EMAIL NotificationType
	SMS   NotificationType
	INAPP NotificationType
}{
	EMAIL: "email",
	SMS:   "sms",
	INAPP: "inapp",
}
