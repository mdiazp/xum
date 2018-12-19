package controllers

// Permission ...
type Permission string

// Permissions ...
type Permissions []Permission

const (
	// PermissionSession ...
	PermissionSession Permission = "Account"

	// PermissionCreateUser ...
	PermissionCreateUser Permission = "CreateUser"
	// PermissionRetrieveUser ...
	PermissionRetrieveUser Permission = "RetrieveUser"
	// PermissionUpdateUser ...
	PermissionUpdateUser Permission = "UpdateUser"
	// PermissionDeleteUser ...
	PermissionDeleteUser Permission = "DeleteUser"

	// PermissionCreateXUser ...
	PermissionCreateXUser Permission = "CreateXUser"
	// PermissionRetrieveXUser ...
	PermissionRetrieveXUser Permission = "RetrieveXUser"
	// PermissionUpdateXUser ...
	PermissionUpdateXUser Permission = "UpdateXUser"
	// PermissionDeleteXUser ...
	PermissionDeleteXUser Permission = "DeleteXUser"

	// PermissionCreateXGroup ...
	PermissionCreateXGroup Permission = "CreateXGroup"
	// PermissionRetrieveXGroup ...
	PermissionRetrieveXGroup Permission = "RetrieveXGroup"
	// PermissionUpdateXGroup ...
	PermissionUpdateXGroup Permission = "UpdateXGroup"
	// PermissionDeleteXGroup ...
	PermissionDeleteXGroup Permission = "DeleteXGroup"
)

func addSessionPerms(ps *Permissions) {
	*ps = append(*ps, PermissionSession)
}
func addUserPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateUser)
	*ps = append(*ps, PermissionRetrieveUser)
	*ps = append(*ps, PermissionUpdateUser)
	*ps = append(*ps, PermissionDeleteUser)
}
func addXUserPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateXUser)
	*ps = append(*ps, PermissionRetrieveXUser)
	*ps = append(*ps, PermissionUpdateXUser)
	*ps = append(*ps, PermissionDeleteXUser)
}
func addXGroupPerms(ps *Permissions) {
	*ps = append(*ps, PermissionCreateXGroup)
	*ps = append(*ps, PermissionRetrieveXGroup)
	*ps = append(*ps, PermissionUpdateXGroup)
	*ps = append(*ps, PermissionDeleteXGroup)
}
func readOnly(ps *Permissions) {
	*ps = append(*ps, PermissionRetrieveUser)
	*ps = append(*ps, PermissionRetrieveXUser)
	*ps = append(*ps, PermissionRetrieveXGroup)
}
func all(ps *Permissions) {
	addSessionPerms(ps)
	addUserPerms(ps)
	addXUserPerms(ps)
	addXGroupPerms(ps)
}

// Rol ...
type Rol string

const (
	// RolAdmin ...
	RolAdmin = "Admin"
	// RolUser ...
	RolUser = "User"
	// RolReadOnly ...
	RolReadOnly = "ReadOnly"
)

// GetPermissions ...
func GetPermissions(rol Rol) Permissions {
	ps := Permissions{}
	switch rol {
	case RolAdmin:
		all(&ps)
	case RolUser:
		ps = append(ps, PermissionSession)
	case RolReadOnly:
		readOnly(&ps)
	}
	return ps
}
