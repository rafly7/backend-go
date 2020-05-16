package models

type Permission struct {
	Base
	PermissionName string `gorm:"type:varchar(8);not null" json:"permissionName"`
}

type Permissions []Permission
