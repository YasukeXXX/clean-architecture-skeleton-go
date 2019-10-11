package registries

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgres() Postgres {
	db, err := RetryConnection(func() (db *gorm.DB, err error) {
		db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
		return
	}, 10)
	if err != nil {
		log.Fatal(err)
	}
	return Postgres{db}
}

type connect func() (*gorm.DB, error)

func RetryConnection(fnc connect, retry int) (db *gorm.DB, err error) {
	for i := 0; i < retry; i++ {
		db, err = fnc()
		if err != nil {
			log.Print(err)
			time.Sleep(10 * time.Second)
			continue
		}
		return
	}
	return db, fmt.Errorf("Failed connect to db")
}
