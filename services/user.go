package services

import (
	"encoding/json"
	"fmt"

	"github.com/rafly7/backend/configs"
	"github.com/rafly7/backend/models"
)

func (c *Connection) GetAllUser(u *models.Users) *models.Users {
	if exists := c.Db.Find(&u).RecordNotFound(); exists {
		return nil
	}
	return u
}

func (c *Connection) GetUserById(id string, u *models.User) *models.User {
	if c.Db.Where("id=?", id).First(&u).RecordNotFound() {
		return nil
	}
	return u
}

func (c *Connection) AddUser(payload []byte, u *models.User) *models.User {
	if err := json.Unmarshal(payload, &u); err != nil {
		return nil
	}
	if len(u.Name) < 3 || len(u.Email) < 6 || len(u.Password) < 6 {
		return nil
	}
	*u = models.User{Name: u.Name, Email: u.Email, Password: configs.HashPassword(u.Password)}
	c.Db.NewRecord(*u)
	c.Db.Create(&u)
	c.Db.NewRecord(*u)
	fmt.Printf("Name: %s, Email: %s, Password: %s", u.Name, u.Email, u.Password)
	return u
}

func (c *Connection) UpdateUser(payload []byte, u *models.User) *models.User {
	if err := json.Unmarshal(payload, &u); err != nil {
		return nil
	}
	if len(u.Name) < 3 || len(u.Email) < 6 || len(u.Password) < 6 {
		return nil
	}
	c.Db.First(&u, u.Id)
	*&u.Name = u.Name
	*&u.Email = u.Email
	fmt.Println(u.Password)
	*&u.Password = configs.HashPassword(u.Password)
	c.Db.Save(&u)
	return u
}

func (c *Connection) DeleteUser(id string, u *models.User) int64 {
	// c.Db.Where("id=?", id).Delete(&u).RowsAffected
	return c.Db.Where("id=?", id).Delete(&u).RowsAffected
}
