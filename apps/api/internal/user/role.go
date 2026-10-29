package user

type UserRole string

const (
	UserRoleOwner  UserRole = "owner"
	UserRoleAdmin  UserRole = "admin"
	UserRoleMember UserRole = "member"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleOwner, UserRoleAdmin, UserRoleMember:
		return true
	}
	return false
}
