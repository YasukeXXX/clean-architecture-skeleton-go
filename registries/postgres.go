package registries

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgres() Postgres {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=skeleton sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return Postgres{db}
}
