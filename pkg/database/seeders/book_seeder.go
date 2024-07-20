package seeders

import (
	"blog/pkg/console"
	"blog/pkg/database/factories"
	"blog/pkg/logger"
	"blog/pkg/seed"
	"fmt"
	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedBooksTable", func(db *gorm.DB) {

		books := factories.MakeBooks(15)

		result := db.Table("books").Create(&books)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded",
			result.Statement.Table, result.RowsAffected))
	})
}
