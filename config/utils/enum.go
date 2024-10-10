package utils

type UserRole string

var EnumUserRole = struct {
	USER  UserRole
	ADMIN UserRole
}{
	USER:  "user",
	ADMIN: "admin",
}
