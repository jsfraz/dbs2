package models

type CreateUser struct {
	Register
	Role Role `json:"role" validate:"oneof=databaseManager reviewApprover,required"`
}
