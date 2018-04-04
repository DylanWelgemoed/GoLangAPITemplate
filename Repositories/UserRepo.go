package repositories

import (
    "time"
	"fmt"
	"github.com/DylanWelgemoed/GoLangAPITemplate/structs"
)

var currentId int

var UserList structs.Users

// Give us some seed data
func init() {
    CreateUser(structs.User{Name: "Dylan Welgemoed", Email: "Test@Test.com"})
    CreateUser(structs.User{Name: "Robert Bogle", Email: "Test@Test.com"})
}

func GetUsers() structs.Users {
    return UserList
}

func FindUser(id int) structs.User {
    for _, u := range UserList {
        if u.Id == id {
            return u
        }
    }
    // return empty User if not found
    return structs.User{}
}

func CreateUser(user structs.User) structs.User {
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