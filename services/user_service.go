package services

import (
    "time"
	"fmt"
	"github.com/DylanWelgemoed/GoLangAPITemplate/models"
)

var currentId int

var UserList models.Users

// Give us some seed data
func init() {
    CreateUser(models.User{Name: "Dylan Welgemoed", Email: "Test@Test.com"})
    CreateUser(models.User{Name: "Robert Bogle", Email: "Test@Test.com"})
}

func GetUsers() models.Users {
    return UserList
}

func FindUser(id int) models.User {
    for _, u := range UserList {
        if u.Id == id {
            return u
        }
    }
    // return empty User if not found
    return models.User{}
}

func CreateUser(user models.User) models.User {
    currentId += 1
    user.Id = currentId
    user.DateCreated = time.Now()
    UserList = append(UserList, user)
    return user
}

func DestroyUser(id int) error {
    for i, u := range UserList {
        if u.Id == id {
            UserList = append(UserList[:i], UserList[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find User with id of %d to delete", id)
}