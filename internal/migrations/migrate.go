package migrations

import (
	"github.com/KadirOzerOzturk/url-shortener/app/entities"
	"github.com/KadirOzerOzturk/url-shortener/internal/database"
)

func init() {
	database.Connection().AutoMigrate(&entities.Url{})
	database.Connection().AutoMigrate(&entities.Log{})
	database.Connection().AutoMigrate(&entities.Mail{})
	database.Connection().AutoMigrate(&entities.User{})
}
