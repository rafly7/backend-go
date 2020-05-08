package services

import (
	"encoding/json"
	"fmt"

	"github.com/rafly7/backend/configs"
	"github.com/rafly7/backend/models"
)

func (c *Connection) GetAllUser(u *models.Users) *models.Users {
	c.Db.Find(&u)
	return u
}

func (c *Connection) GetUserById(id string, u *models.User) *models.User {
	exists := c.Db.Where("id=?", id).First(&u).RecordNotFound()
	if exists {
		return nil
	}
	return u
	// var exists bool = c.Db.Where("id=?", id).First(&u).RecordNotFound()
	// if exists {
	// 	return u, exists
	// } else {
	// 	return u, exists
	// }
}

func (c *Connection) AddUser(payload []byte, u *models.User) *models.User {
	json.Unmarshal(payload, &u)
	*u = models.User{Name: u.Name, Email: u.Email, Password: configs.HashPassword(u.Password)}
	c.Db.NewRecord(*u)
	c.Db.Create(&u)
	c.Db.NewRecord(*u)
	fmt.Printf("Name: %s, Email: %s, Password: %s", u.Name, u.Email, u.Password)
	return u
}

func (c *Connection) UpdateUser(payload []byte, u *models.User) *models.User {
	json.Unmarshal(payload, &u)
	c.Db.First(&u, u.Id)
	*&u.Name = u.Name
	*&u.Email = u.Email
	*&u.Password = configs.HashPassword(u.Password)
	c.Db.Save(&u)
	return u
}

func (c *Connection) DeleteUser(id string, u *models.User) *models.User {
	c.Db.Where("id=?", id).Delete(&u)
	return u
}
