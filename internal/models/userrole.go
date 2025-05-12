package models

type UserRole int8

const (
	RoleUser UserRole = iota
	RoleAdmin
)

func (r UserRole) String() string {
	switch r {
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}
