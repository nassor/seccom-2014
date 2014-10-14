package mongo

import (
	"mongo/models"

	"gopkg.in/mgo.v2"
)

func AddUser(c *mgo.Collection, user *models.User) error {
	err := c.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func ListUsers(c *mgo.Collection) ([]models.User, error) {
	var users []models.User

	query := c.Find(nil).Sort("name")
	err := query.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
