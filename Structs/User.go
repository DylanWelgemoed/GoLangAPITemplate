package structs

import "time"

type User struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Email      string    `json:"email"`
    DateCreated       time.Time `json:"dateCreated"`
}

type Users []User