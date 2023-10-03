package userservice

import (
    contact_service "contactapp_api/components/contacts/contact_service"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "log"
    "errors"
)

type User struct {
    Id       string
    Name     string
    Email    string
    Password string
    Contacts []*contact_service.Contact
}

var users = []*User{}

func handlePanic(err interface{}) {
    log.Printf("Panic: %v\n", err)
}

func CreateAdmin(Name, Email, Password string) *User {
    defer func() {
        if r := recover(); r != nil {
            handlePanic(r)
        }
    }()

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), 7)
    if err != nil {
        panic(err)
    }

    user := &User{
        Id:       uuid.NewString(),
        Name:     Name,
        Email: Email,
        Password: string(hashedPassword),
        Contacts: []*contact_service.Contact{},
    }

    users = append(users, user)
    return user
}

func CreateUser(Name, Email, Password string) *User {
    defer func() {
        if r := recover(); r != nil {
            handlePanic(r)
        }
    }()

    hashedPass, err := bcrypt.GenerateFromPassword([]byte(Password), 7)
    if err != nil {
        panic(err)
    }

    user := &User{
        Id:       uuid.NewString(),
        Name:     Name,
        Email: Email,
        Password: string(hashedPass),
        Contacts: []*contact_service.Contact{},
    }

    users = append(users, user)
    return user
}

func InsertContact(user *User, contact *contact_service.Contact) {
    user.Contacts = append(user.Contacts, contact)
}

func GetAllUsers() []User {
    var allUsers = []User{}
    for _, user := range users {
        allUsers = append(allUsers, *user)
    }

    return allUsers
}

func GetUserById(id string) (*User, bool) {
    for _, user := range users {
        if user.Id == id {
            return user, true
        }
    }

    return nil, false
}

func UpdateUserById(body *User, user *User) {
    if body.Name != "" && body.Name != user.Name {
        user.Name = body.Name
    }
}

func DeleteUserByID(id string) {
    defer func() {
        if r := recover(); r != nil {
            handlePanic(r)
        }
    }()

    for i, j := range users {
        if j.Id == id {
            users = append(users[:i], users[i+1:]...)
            return
        }
    }
    panic("User not found")
}


func GetUserByEmail(Email string) (*User,error){

    for _,j := range users {
        if j.Email == Email {
            return j,nil
        }
    }
    return nil,errors.New("user does not exist")

}

