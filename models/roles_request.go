package models

type RolesRequest struct {
	Roles []Role `query:"roles" validate:"required,unique"`
}
