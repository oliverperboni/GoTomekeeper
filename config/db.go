// config/config.go
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "user=postgres password=oliver dbname=bookShelf port=5432 sslmode=disable"
	dsn := "host=localhost user=postgres password=oliver dbname=bookShelf port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//&schemas.ListBook{}, &schemas.List{}, &schemas.ReadingStatistic{}, &schemas.Review{}, &schemas.User{}, &schemas.UserBook{})
	// db.AutoMigrate(&schemas.User{}, &schemas.Book{}, &schemas.UserBook{}, &schemas.List{}, &schemas.ListBook{}, &schemas.Review{}, &schemas.ReadingStatistic{})
	// dummyUser := schemas.User{
	// 	Username:     "dummyUser",
	// 	Email:        "dummy@example.com",
	// 	PasswordHash: "hashedpassword",
	// 	CreatedAt:    time.Now(),
	// 	UpdatedAt:    time.Now(),
	// }
	// db.Create(&dummyUser)
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
