package model

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Base struct {
	ID        uint       `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time  `gorm:"autoUpdateTime:milli" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt *time.Time `gorm:"default:null" json:"deleted_at,omitempty"`
}

type User struct {
	Base
	Name    string   `gorm:"unique" json:"name"`
	Surname string   `json:"surname"`
	Age     int      `json:"age"`
	Status  bool     `json:"status"`
	Address *Address `json:"address,omitempty"`
}

func (e *User) Disable() {
	e.Status = false
}

func (p *User) Enable() {
	p.Status = true
}

type Address struct {
	Base
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Address{})
	return db
}
