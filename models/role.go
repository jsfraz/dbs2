package models

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleDbManager Role = "databaseManager"
	RoleReview    Role = "reviewApprover"
	RoleCustomer  Role = "customer"
)
