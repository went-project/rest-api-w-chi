package models

import (
	"time"
	"went-template/internal/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id" swaggerignore:"true"`
	Username  string         `gorm:"unique;not null" json:"username" validate:"required_if=ID 0,min=3,max=50" example:"john doe"`
	Email     string         `gorm:"unique;not null" json:"email" validate:"required_if=ID 0,email" example:"johndoe@example.com"`
	Password  string         `gorm:"not null" json:"password" validate:"required_if=ID 0,min=6" example:"changeme"`
	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}

// HiddenFields specifies fields to omit in JSON responses
var HiddenFields = []string{
	"password",
	"deleted_at",
}

// TableName sets the insert table name for this struct type
func (User) TableName() string {
	return "users"
}

// Validate validates the user fields.
func (u *User) Validate() error {
	err := validator.New().Struct(u)
	if err != nil {
		return err
	}
	return nil
}

// BeforeCreate is a GORM hook that is called before a new record is inserted into the database.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Hash password
	u.Password = utils.Hash(u.Password)
	return
}

// BeforeUpdate is a GORM hook that is called before an existing record is updated in the database.
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Hash password if it has changed
	if tx.Statement.Changed("password") {
		u.Password = utils.Hash(u.Password)
	}
	return
}

// CRUD Operations
func (u *User) FindAll(db *gorm.DB, where map[string]interface{}) ([]User, error) {
	var users []User
	err := db.Select("*").Omit(HiddenFields...).Where(where).Find(&users).Error
	return users, err
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}
