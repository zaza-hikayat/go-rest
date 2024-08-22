package constants

var (
	ROLE_SUPERADMIN = "super_admin"
	ROLE_ADMIN      = "admin"
	ROLE_MEMBER     = "member"
)

var mapRole = map[string]string{
	ROLE_SUPERADMIN: "Super Admin",
	ROLE_ADMIN:      "Admin",
}

func GetRoleName(name string) string {
	if v, ok := mapRole[name]; ok {
		return v
	}

	return ""
}
