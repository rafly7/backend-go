package services

import (
	"encoding/json"

	"github.com/rafly7/backend/models"
)

func (c *Connection) GetAllPermission(u *models.Permissions) *models.Permissions {
	if exists := c.Db.Find(&u).RecordNotFound(); exists {
		return nil
	}
	return u
}

func (c *Connection) GetPermissionById(id string, u *models.Permission) *models.Permission {
	if c.Db.Where("id=?", id).First(&u).RecordNotFound() {
		return nil
	}
	return u
}

func (c *Connection) AddPermission(payload []byte, u *models.Permission) *models.Permission {
	if err := json.Unmarshal(payload, &u); err != nil {
		return nil
	}
	if len(u.PermissionName) < 3 {
		return nil
	}
	*u = models.Permission{PermissionName: u.PermissionName}
	c.Db.NewRecord(*u)
	c.Db.Create(&u)
	c.Db.NewRecord(*u)
	return u
}

func (c *Connection) UpdatePermission(payload []byte, u *models.Permission) *models.Permission {
	if err := json.Unmarshal(payload, &u); err != nil {
		return nil
	}
	if len(u.PermissionName) < 3 {
		return nil
	}
	c.Db.First(&u, u.Id)
	*&u.PermissionName = u.PermissionName
	c.Db.Save(&u)
	return u
}

func (c *Connection) DeletePermission(id string, u *models.Permission) int64 {
	// c.Db.Where("id=?", id).Delete(&u).RowsAffected
	return c.Db.Where("id=?", id).Delete(&u).RowsAffected
}
