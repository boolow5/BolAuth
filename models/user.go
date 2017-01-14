package models

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID         int       `json:"id" gorm:"primary_key"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role       Role      `json:"role"`
	TimeAdded  time.Time `json:"time_added"`
	Active     bool      `json:"active"`
}

func (this *User) String() string {
	if len(this.FirstName) > 0 {
		return fmt.Sprintf("%s %s %s", this.FirstName, this.MiddleName, this.LastName)
	}
	return fmt.Sprintf("%s", this.Username)
}

func (this *User) Add() (bool, error) {
	this.TimeAdded = time.Now()
	result := db.Create(this)
	if result.Error != nil {
		fmt.Println("Add Error:", result.Error)
		return false, result.Error
	}
	return true, nil
}

func GetUserByID(id int) (*User, error) {
	User := User{}
	result := db.Where("id = ?", id).First(&User)
	if result.Error != nil {
		return nil, result.Error
	}
	return &User, nil
}

func GetAllUsers() (*[]User, error) {
	todos := []User{}
	result := db.Find(&todos)
	if result.Error != nil {
		return &[]User{}, result.Error
	}
	return &todos, nil
}

func (this *User) Update() (bool, error) {
	if this.FirstName == "" || this.FirstName == " " {
		return false, errors.New("Empty First name")
	}
	if this.LastName == "" || this.LastName == " " {
		return false, errors.New("Empty Last name")
	}
	result := db.Save(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (this *User) Delete() (bool, error) {
	result := db.Delete(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
