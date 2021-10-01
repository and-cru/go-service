package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID      uuid.UUID `gorm:"unique;primaryKey;type:uuid"`
	Name    string    `gorm:"unique" json:"name"`
	Surname string    `json:"surname"`
	Age     int       `json:"age"`
	Status  bool      `json:"status"`
	Address Address   `json:"address"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid)
}

func (e *User) Disable() {
	e.Status = false
}

func (p *User) Enable() {
	p.Status = true
}

type Address struct {
	ID      uuid.UUID `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	Country string    `json:"country"`
}

func (address *Address) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid)
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Address{})
	return db
}
